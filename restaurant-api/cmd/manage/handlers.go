package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"lightsaid.com/restaurant-app/restaurant-api/internal/model"
	"lightsaid.com/restaurant-app/restaurant-api/internal/security"
)

var tokenExpiredAt = 10 * time.Minute

func (app *application) createAdmin(ctx *gin.Context) {
	var req = new(createAdminRequest)
	if ok := app.bindError(ctx, req); !ok {
		return
	}
	hashedPwd, err := security.GenerateHashPwd(req.Password)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"admin": nil, "err": err.Error(), "msg": "加密出错"})
		return
	}
	admin, err := app.Repo.CreateAdmin(req.Name, req.Phone, hashedPwd)
	if err != nil {
		// TODO: 索引错误处理
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"admin": nil, "err": err.Error(), "msg": "创建失败"})
		return
	}
	admin.Password = nil
	ctx.JSON(http.StatusOK, gin.H{"admin": admin})
}

func (app *application) login(ctx *gin.Context) {
	var req = new(loginRequest)
	if ok := app.bindError(ctx, req); !ok {
		return
	}
	a, err := app.Repo.GetAdminByPhone(req.Phone)
	if err != nil {
		// TODO: 不存在错误处理
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "err": err.Error(), "msg": "登录出错"})
		return
	}

	err = security.CheckPassword(req.Password, *a.Password)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "err": err.Error(), "msg": "密码不正确"})
		return
	}
	timeout := tokenExpiredAt // 默认 10 分钟
	if d, err := time.ParseDuration(os.Getenv("TOKEN_TIMEOUT")); err != nil {
		zap.S().Error("time.ParseDuration error -> ", err)
	} else {
		timeout = d
	}
	token, assPayload, err := app.jwtMaker.CreateToken(a.ID.Hex(), timeout)
	if err != nil {
		// TODO: 详细错误处理
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "err": err.Error(), "msg": "服务内部出错"})
		return
	}

	ref_timeout := tokenExpiredAt
	if t, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TIMEOUT")); err != nil {
		zap.S().Error("time.ParseDuration error -> ", err)
	} else {
		ref_timeout = t
	}
	refreshToken, payload, err := app.jwtMaker.CreateToken(a.ID.Hex(), ref_timeout)
	if err != nil {
		// TODO: 详细错误处理
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "err": err.Error(), "msg": "服务内部出错"})
		return
	}

	idStr, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "err": err.Error(), "msg": "objectID 转换出错"})
		return
	}
	// 创建 session
	s := model.Session{
		ID:           idStr,
		UserID:       a.ID.Hex(),
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		CreatedAt:    time.Now(),
		RefreshToken: refreshToken,
		ExpiredAt:    payload.ExpiredAt,
	}
	ss, err := app.Repo.CreateRefreshToken(s)
	if err != nil {
		// TODO: 详细错误处理
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "err": err.Error(), "msg": "服务内部出错"})
		return
	}
	a.Password = nil
	data := loginResponse{
		Admin:                 a,
		Token:                 token,
		SessionID:             ss.ID.Hex(),
		ExpiredAt:             assPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: ss.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data, "err": nil, "msg": "成功"})
}

func (app *application) createShop(ctx *gin.Context) {

	req := createShopRequest{}
	if ok := app.bindError(ctx, &req); !ok {
		return
	}
	shop, err := app.Repo.CreateShop(req.Name, req.Logo)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"shop": nil, "err": err.Error(), "msg": "创建失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"shop": shop})
}

func (app *application) getShops(ctx *gin.Context) {
	list, err := app.Repo.GetShops()
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"list": nil, "err": err.Error(), "msg": "获取列表出错"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"list": list, "err": nil, "msg": "成功"})
}

func (app *application) renewAccessToken(ctx *gin.Context) {
	var req = new(renewAccessTokenRequest)
	if ok := app.bindError(ctx, req); !ok {
		return
	}

	resp := renewAccessTokenResponse{}

	payload, err := app.jwtMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"data": nil, "err": err.Error(), "msg": err.Error()})

		return
	}
	oid, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"data": nil, "err": err.Error(), "msg": "刷新token错误"})

		return
	}

	fmt.Println(">>> oid ", oid)
	session, err := app.Repo.GetRefreshToken(oid)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"data": nil, "err": err.Error(), "msg": "刷新token错误"})

		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("session 已被篡改")
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"data": nil, "err": err.Error(), "msg": "刷新token错误"})
		return
	}

	if session.UserID != payload.UserID || session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("session 不合法")
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"data": nil, "err": err.Error(), "msg": err.Error()})
		return
	}

	if time.Now().After(session.ExpiredAt) {
		err := fmt.Errorf("session 已过期, 请重写登录")
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"data": nil, "err": err.Error(), "msg": err.Error()})
		return
	}

	// 可以再严格判断 ip 和 useragent
	if session.ClientIP != ctx.ClientIP() || session.UserAgent != ctx.Request.UserAgent() {
		err := fmt.Errorf("session 不合法")
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"data": nil, "err": err.Error(), "msg": err.Error()})
		return
	}

	timeout := tokenExpiredAt // 默认 10 分钟
	if d, err := time.ParseDuration(os.Getenv("TOKEN_TIMEOUT")); err != nil {
		zap.S().Error("time.ParseDuration error -> ", err)
	} else {
		timeout = d
	}

	// 验证refresh_token是否合法有效
	token, p, err := app.jwtMaker.CreateToken(session.UserID, timeout)
	if err != nil {
		// TODO: 详细错误处理
		zap.S().Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": nil, "err": err.Error(), "msg": "服务内部出错"})
		return
	}

	resp.ExpiredAt = p.ExpiredAt
	resp.Token = token
	ctx.JSON(http.StatusOK, gin.H{"data": resp, "err": nil, "msg": "成功"})
}
