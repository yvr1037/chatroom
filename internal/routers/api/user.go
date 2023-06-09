package api

import (
	"chatroom/internal/request"
	"chatroom/internal/service"
	"chatroom/pkg/errcode"
	"chatroom/pkg/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Register(c *gin.Context) {
	param := request.UserRegisterRequest{}
	r := response.NewResponse(c)
	if c.ShouldBindJSON(&param) != nil {
		r.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UserRegister(&param)
	if err != nil {
		r.ToErrorResponse(err)
		return 
	}
	r.ToResponse(nil)
}

func (u User) Login(c *gin.Context) {
	param := request.UserLoginRequest{}
	response := response.NewResponse(c)
	if c.ShouldBindJSON(&param) != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	user,err := svc.UserLogin(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return 
	}
	response.ToResponse(user)
}