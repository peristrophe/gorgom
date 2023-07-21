package util

import (
	"fmt"
	"gorgom/internal/setting"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT string

func NewJWT(userID string) *JWT {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * time.Duration(setting.TOKEN_EXPIRE)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(setting.TOKEN_SECRET_KEY))
	if err != nil {
		return nil
	}
	jwt := JWT(tokenStr)
	return &jwt
}

func (t *JWT) Parse() (*jwt.Token, error) {
	// contain verification also
	tokenStr := string(*t)
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(setting.TOKEN_SECRET_KEY), nil
	})
}

func (t *JWT) WhoAmI() string {
	token, err := t.Parse()
	if err != nil {
		return ""
	}
	claims := token.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%s", claims["userID"])
}
