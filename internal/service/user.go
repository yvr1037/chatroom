package service

import (
	"chatroom/internal/request"
	"chatroom/pkg/auth"
	"chatroom/pkg/errcode"
)

type LoginRespondContent struct {
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) UserRegister(param *request.UserRegisterRequest) *errcode.Error {
	return svc.dao.UserRegister(param.UserName, param.Nickname, param.Password)
}

func (svc *Service) UserLogin(param *request.UserLoginRequest) (*LoginRespondContent, *errcode.Error) {
	user, err := svc.dao.UserLogin(param.UserName, param.Password)
	if err != nil {
		return nil, err
	}
	ID := user.ID
	token, _ := auth.GenerateToken(ID)
	return &LoginRespondContent{
		UserID: ID,
		Token:  token,
	}, nil
}
