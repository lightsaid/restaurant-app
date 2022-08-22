package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func handleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// bindArgError 公共处理请求JSON参数，如果返回 true 则代表解析参数成功
func (app *application) bindError(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		errFileds, ok := err.(validator.ValidationErrors)
		if ok {
			// 返回来的是 map[string]string 类型
			errinfo := errFileds.Translate(app.trans)
			// 提起一个错误字符串
			errMsgs := []string{}
			for _, val := range errinfo {
				errMsgs = append(errMsgs, val)
			}
			if len(errMsgs) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"msg": errMsgs[0], "err": errinfo})
				return false
			}
			c.JSON(http.StatusBadRequest, gin.H{"msg": errinfo, "err": errinfo})
			return false
		}
		// 其他的一些错误
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不能为空", "err": err})
			return false
		}
		if strings.Contains(err.Error(), "cannot unmarshal") {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "参数类型不匹配", "err": err})
			return false
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": "解析参数错误", "err": err})
		return false
	}
	return true
}
