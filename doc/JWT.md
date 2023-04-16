```go
type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}
```
这个结构体定义了一个JWT设置,其中包含以下字段:
- Secret:JWT的密钥,用于签名和验证;
- Issuer:发布JWT的实体,通常是一个字符串标识。
- Expire:JWT的过期时间,表示从创建时间开始多长时间之后JWT将过期;它是一个time.Duration类型的值;
这个结构体可以用来存储应用程序中使用的JWT设置,以便在需要时轻松地引用它们