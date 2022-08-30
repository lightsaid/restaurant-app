package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"lightsaid.com/restaurant-app/restaurant-api/internal/security"
)

// 定义常量
const (
	// 存 token header 的 key
	AuthHeaderKey = "Authorization"

	// token 字符串前缀
	AuthHeaderPrefix = "bearer"

	// 存 gin.Context token playload key
	AutPayloadKey = "auth_pyload"
)

// authMiddleware 验证是否登录
func authMiddleware(jwtMaker *security.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		authHeader := c.GetHeader(AuthHeaderKey)
		if len(authHeader) == 0 {
			err := errors.New("权限验证无效,Header不存在token信息")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("请先登录", err))
			return
		}
		// 以空格分割字token "Bearer eyJhbGciOiJI..."
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("token 验证无效")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error(), err))
			return
		}

		// 校验 token 头是否是 Bearer
		prefix := strings.ToLower(fields[0])
		if prefix != AuthHeaderPrefix {
			err := fmt.Errorf("token 头不正确")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error(), err))
			return
		}

		// 获取 token 值
		accessToken := fields[1]
		payload, err := jwtMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error(), err))
			return
		}
		zap.S().Info("gin.Context: ", *payload)
		c.Set(AutPayloadKey, *payload)
		c.Next()
	}
}

// rateLimit 请求限流控制，根据请求IP地址限流
func rateLimit(rate int) gin.HandlerFunc {
	var smap sync.Map
	return func(c *gin.Context) {
		// 获取客户端请求IP
		ip := c.ClientIP()
		l, ok := smap.Load(ip)
		if !ok {
			l = ratelimit.New(rate)
		}
		lm, ok := l.(ratelimit.Limiter)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "中间件类型断言错误limiter", "msg": "服务内部错误"})

			// c.Abort()
			return
		}
		lm.Take()
		smap.Store(ip, lm)
		c.Next()
	}
}

// requestTimeout 请求超时处理
func requestTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// accessLogWiter 访问日志结构体
type accessLogWiter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 实现 gin.ResponseWriter 的 Write 方法
func (l accessLogWiter) Write(b []byte) (int, error) {
	if n, err := l.body.Write(b); err != nil {
		return n, err
	}
	return l.ResponseWriter.Write(b)
}

func accessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &accessLogWiter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		c.Writer = bodyWriter
		startTime := time.Now().Format("2006/01/02 15:04:05")
		c.Next()
		endTime := time.Now().Format("2006/01/02 15:04:05")

		// query := c.Request.URL.Query().Encode()

		// var sb []byte
		// _, _ = c.Request.Body.Read(sb)
		inParam := make(map[string]interface{})

		data := struct{}{}
		dec := json.NewDecoder(c.Request.Body)
		_ = dec.Decode(&data)

		postval := c.Request.PostForm.Encode()

		// inParam["Query"] = query // 因为 uri 包含 query
		inParam["Body"] = data //string(sb) // TODO: 获取不到数据
		inParam["PostForm"] = postval

		zap.S().Infof("\nACCESSLOG: Time: %s ~ %s, Method: %s, Uri: %s, IN: %v, OUT: %v\n",
			startTime, endTime, c.Request.Method, c.Request.URL, inParam, bodyWriter.body.String())
	}
}
