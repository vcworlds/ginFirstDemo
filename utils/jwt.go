package utils

import (
	"gin-Vue/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type Claims struct {
	Id uint
	jwt.StandardClaims
}

// jwt的密钥
var jwtKey = []byte("ASDFGHJKLZXCVBNMasdfghjklxcvbnm")

// ReleaseToken jwt的生成
func ReleaseToken(user *models.User) string {
	// 过期时间
	expiredTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		Id: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Syy",
			Subject:   "user_jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Print("生成token失败", err)
		return ""
	}
	return tokenString

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})
	return token, claims, nil
}
