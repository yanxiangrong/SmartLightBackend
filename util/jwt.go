package util

import (
	"SmartLightBackend/pkg/logging"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

var jwtSecret = fmt.Sprintf("%x", rand.Int())

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		logging.Error("Signed String Fail ", err.Error())
	}
	return token
}

func ParseToken(token string) Claims {
	_, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		logging.Info("Signed String Fail ", err.Error())
	}

	return Claims{}
}
