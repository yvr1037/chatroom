package auth

import (
	"chatroom/global"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSettings.Secret)
}

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