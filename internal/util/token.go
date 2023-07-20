package util

import (
	"fmt"
	"gorgom/internal/setting"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(setting.TOKEN_EXPIRE)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(setting.TOKEN_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(setting.TOKEN_SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
