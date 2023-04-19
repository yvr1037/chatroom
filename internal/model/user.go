package model

import (
	"chatroom/pkg/errcode"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `json:"user_id" gorm:"autoIncrement"`
	Username string `json:"user_name"`
	NickName string `json:"nick_name"`
	// Password  string `json:"password"`
	Password  string `json:"-"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (u User) Create(db *gorm.DB) *errcode.Error {
	var user User
	err := db.Where("user_name = ?", u.Username).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		return errcode.ErrorDuplicatedUserName
	}

	// return errcode.Convert(db.Create(&user).Error)
	err = db.Create(&u).Error
	if err != nil {
		return errcode.Convert(err)
	}
	return nil
}

func (u User) Get(db *gorm.DB) (*User, *errcode.Error) {
	var user User
	var err error
	if u.ID == 0 {
		err := db.Where("user_name = ?", u.Username).First(&user).Error
		if err == gorm.ErrRecordNotFound {
			return &user, errcode.ErrorUserNameNotFound
		}
	} else {
		err = db.Where("id = ?", u.ID).First(&user).Error
		if err == gorm.ErrRecordNotFound {
			return &user, errcode.ErrorUserNameNotFound
		}
	}
	// return user,errcode.Convert(err)
	if err != nil {
		return &user, errcode.Convert(err)
	}
	return &user, nil
}
