package request

type UserRegisterRequest struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Nickname string `form:"nick_name" json:"nick_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"password"`
}

type UserLoginRequest struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}