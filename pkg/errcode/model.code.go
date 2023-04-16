package errcode

var (
	ErrorDuplicatedUserName = NewError(20010001, "UserName Repeat")
	ErrorUserNameNotFound   = NewError(20010002, "account notfound")
	ErrorUserIDNotFound     = NewError(20010003, "用户ID不存在")
	ErrorPassword           = NewError(20010004, "账号或密码错误")
	ErrorRegisterFailure    = NewError(20010005, "注册失败")
	ErrorLoginFailure       = NewError(20010006, "登录失败")
	// ErrorPassword = NewError(20010003,"account or password incorrect")
	// ErrorRegisterFail = NewError(20010004,"register fail")
	// ErrorLoginFail = NewError(20010005,"login fail")

	ErrorSendMessageFail = NewError(20020001, "send msg fail")
	ErrorGetMessageFail  = NewError(20020002, "get msg fail")
)
