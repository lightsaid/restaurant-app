package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"lightsaid.com/restaurant-app/restaurant-api/internal/security"
)

func (app *application) createAdmin(ctx *gin.Context) {
	var req = new(createAdminRequest)
	if ok := app.bindError(ctx, req); !ok {
		return
	}
	hashedPwd, err := security.GenerateHashPwd(req.Password)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"admin": nil, "err": err, "msg": "加密出错"})
		return
	}
	admin, err := app.Repo.CreateAdmin(req.Name, req.Phone, hashedPwd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"admin": nil, "err": err, "msg": "创建失败"})
		return
	}
	admin.Password = nil
	ctx.JSON(http.StatusOK, gin.H{"admin": admin})
}

func (app *application) createShop(ctx *gin.Context) {
	shop, err := app.Repo.CreateShop("小猪猪", "http://localhost:4000/files/logo.png")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"shop": nil, "err": err, "msg": "创建失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"shop": shop})
}

func (app *application) getShop(ctx *gin.Context) {

}

func (app *application) getShops(ctx *gin.Context) {

}

func (app *application) updateShop(ctx *gin.Context) {

}

func (app *application) deleteShop(ctx *gin.Context) {

}
