package common

import (
	"gin01/app/v1/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// jwt鉴权
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}
func ReleaseToken (user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: uint(user.ID),
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "hackbar.tech",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "",err

	}
	return tokenString,nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, Claims , func(token *jwt.Token) (i interface{},err error) {
		return jwtKey, nil
	})

	return token, Claims, err
}

func GetRequestIP(c *gin.Context)string{
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}