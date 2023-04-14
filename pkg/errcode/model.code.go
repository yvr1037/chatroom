package errcode

var (
	ErrorDuplicatedUserName = NewError(20010001,"UserName Repeat")
	ErrorUserNameNotFound = NewError(20010002,"account notfound")
	ErrorPassword = NewError(20010003,"account or password incorrect")
	ErrorRegisterFail = NewError(20010004,"register fail")
	ErrorLoginFail = NewError(20010005,"login fail")

	ErrorSendMessageFail = NewError(20020001,"send msg fail")
	ErrorGetMessageFail = NewError(20020002,"get msg fail")
)