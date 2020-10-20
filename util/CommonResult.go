package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResult(c *gin.Context,value interface{}){
	c.JSON(http.StatusOK,gin.H{
		"code" : "ok",
		"data" : value,
	})
}
