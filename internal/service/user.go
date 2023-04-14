package service

import (
	"chatroom/internal/model"
	"chatroom/internal/request"
	"chatroom/pkg/errcode"
)

func (svc *Service) UserRegister(param *request.UserRegisterRequest) (*model.User,*errcode.Error) {
	return svc.dao.UserRegister(param.UserName,param.Nickname,param.Password)
}


func (svc *Service) UserLogin(param *request.UserLoginRequest) (*model.User,*errcode.Error) {
	return svc.dao.UserLogin(param.UserName,param.Password)
}