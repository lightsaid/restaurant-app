package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	ga := r.Group("/v1/admin")

	ga.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	ga.POST("/shop", app.createShop)

	return r
}
