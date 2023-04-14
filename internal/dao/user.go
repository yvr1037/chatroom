package dao

import (
	"chatroom/internal/model"
	"chatroom/pkg/errcode"
)

func(d *Dao) UserRegister(username,nickname,password string) (*model.User,*errcode.Error) {
	user := model.User{Username: username, NickName: nickname, Password: password}
	err := user.Create(d.engine)
	if err != nil {
		return nil,err
	}
	return d.UserLogin(username,password)
}

func (d *Dao) UserLogin(username,password string) (*model.User,*errcode.Error) {
	user,err := model.User{Username: username}.Get(d.engine)
	if err != nil {
		return nil,err
	}
	if user.Password != password {
		return nil,errcode.ErrorPassword
	}
	return &model.User{
		ID:	user.ID,
		Username: user.Username,
		NickName: user.NickName,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	},nil
}

