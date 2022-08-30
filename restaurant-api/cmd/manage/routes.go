package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()

	r.Use(
		rateLimit(1),                  // 秒控10个
		requestTimeout(3*time.Second), // 1 分钟超时
		accessLog(),
	)

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	ga := r.Group("/v1/admin")
	{
		// 管理员
		ga.POST("/manager", app.createAdmin)
		ga.POST("/login", app.login)

		// 店铺
		// ga.POST("/shop", authMiddleware(app.jwtMaker), app.createShop)
		auth := ga.Group("").Use(authMiddleware(app.jwtMaker))
		{
			auth.POST("/shop", app.createShop)
			auth.GET("/shop", app.getShops)

			auth.POST("/refresh_token", app.renewAccessToken)
		}
	}

	return r
}
