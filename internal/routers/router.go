package routers

import (
	"chatroom/global"
	"chatroom/internal/routers/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSettings.Runmode == "debug" {
		r.Use(gin.Logger(),gin.Recovery()) 
		//gin.Logger() 用于记录每个请求的信息,而gin.Recovery()用于从任何panic中恢复,并向客户端发送一个500错误响应;
	}

	// api := r.Group("/api")
	// {
	// 	api.GET("/",func(c *gin.Context){
	// 		c.JSON(http.StatusOK,gin.H{"msg":"test"})
	// 	})
	// }

	user := api.NewUser()
	message := api.NewMessage()
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{"msg":"test"})
		})

		apiGroup.POST("/register",user.Register)
		apiGroup.POST("/login",user.Login)
		apiGroup.POST("/send",message.Send)
	}
	return r
}
