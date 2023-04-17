package middleware

import (
	"chatroom/pkg/auth"
	"chatroom/pkg/errcode"
	"chatroom/pkg/response"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc { // Gin 中间件函数,用于身份验证和授权;
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		code := errcode.Success
		if token == "" {
			code = errcode.InvalidParams
		}else {
			claims,err := auth.ParseToken(token)
			if err != nil {
				code = errcode.UnauthorizedTokenError
			}else {
				c.Set("UserID",claims.ID)
			}
		}
		if code != errcode.Success {
			r := response.NewResponse(c)
			r.ToErrorResponse(code)
			c.Abort()
			return
		}
		c.Next()
	}
}