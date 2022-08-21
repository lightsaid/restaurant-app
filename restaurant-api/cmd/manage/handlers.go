package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
