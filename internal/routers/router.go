package routers

import (
	"chatroom/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSettings.Runmode == "debug" {
		r.Use(gin.Logger(),gin.Recovery())
	}

	api := r.Group("/api")
	{
		api.GET("/",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{"msg":"test"})
		})
	}
	return r
}
