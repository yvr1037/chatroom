package service

import (
	"chatroom/internal/model"
	"chatroom/internal/request"
	"chatroom/pkg/auth"
	"chatroom/pkg/errcode"
)

type LoginRespondContent struct {
	UserID uint64 `json:"user_id"`
	Nickname string `json:"nick_name"`
	Token string `json:"token"`
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
		Nickname: user.NickName,
		Token:  token,
	}, nil
}

func (svc *Service) UserGet(userID uint64) (*model.User,*errcode.Error){
	return svc.dao.UserGet(userID)
}
