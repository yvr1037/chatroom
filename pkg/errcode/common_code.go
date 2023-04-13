package errcode

var (
	Success                   = NewError(0, "success")
	ServerError               = NewError(10000000, "service internal error")
	InvalidParams             = NewError(10000001, "invalid params")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败,找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败,Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败,Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
)
