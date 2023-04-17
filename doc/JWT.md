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

```go
func GenerateToken(ID uint64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSettings.Expire)
	claims := Claims{
		UserID: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			Issuer:    global.JWTSettings.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetJWTSecret())
	return tokenString, err
}
```
在这个函数中，生成 JWT token 字符串的过程包括以下步骤:获取当前时间nowTime.并计算出 token 的到期时间 expireTime。创建一个 Claims 对象，并将传入的用户 ID 保存在其中的 UserID 属性中。设置该 Claims 对象的标准声明（RegisteredClaims），其中包括：到期时间 (ExpiresAt)、颁发时间 (IssuedAt)、生效时间 (NotBefore) 和签发者 (Issuer)。这些信息将被用来验证和解析 token。使用 jwt.NewWithClaims() 函数创建一个新的 token 对象 tokenClaims，并指定加密算法为 HS256。调用 tokenClaims.SignedString() 方法生成 token 字符串 token，其中使用了通过调用 GetJWTSecret() 函数获取的配置文件中的 JWT 密钥进行签名。最终，该函数将返回生成的 JWT token 字符串以及可能出现的错误信息。

```go
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
```
这个函数用于解析 JWT token 字符串并返回其中包含的声明信息或者错误信息。函数的具体实现如下：使用 jwt.ParseWithClaims() 函数解析传入的 JWT token 字符串，并将解析后的 token 保存到变量 token 中。同时，该函数还需要传入一个回调函数作为第三个参数，用于验证 token 的有效性。如果解析出错，则返回 nil 和错误信息。如果解析成功，则判断 token 是否为 nil，并尝试将其转换为 Claims 结构体类型。如果转换成功并且 token 的有效性也被验证通过，则返回包含声明信息的 Claims 对象和 nil 错误信息。如果上述条件都不满足，则返回 nil 和错误信息。需要注意的是，在回调函数中，使用 GetJWTSecret() 函数获取了配置文件中的 JWT 密钥，以便用于验证 token 的有效性。此外，在解析 token 时，还需要使用 &Claims{} 将一个空的 Claims 结构体指针传递给 jwt.ParseWithClaims() 函数，以便在解析过程中将 token 中的声明信息存储到其中。