package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-mvc/app/model"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//发布token
func ReleaseToken(user model.User) (string, error) {

	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin.mvc",
			Subject:   "gin mvc token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析 token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {

	Claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, Claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, Claims, err
}
