package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var mySecret = []byte("就这样被你征服")

type MyClaims struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userId uint64, username string) (string, error) {
	c := MyClaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer:    "BlueBell",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//加盐
	return token.SignedString(mySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
