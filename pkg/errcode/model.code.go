package errcode

var (
	ErrorDuplicatedUserName = NewError(20010001,"UserName Repeat")
	ErrorPassword = NewError(20010002,"account or password incorrect")

	ErrorSendMessageFail = NewError(20020001,"send msd fail")
	ErrorGetMessageFail = NewError(20020002,"get msg fail")
)