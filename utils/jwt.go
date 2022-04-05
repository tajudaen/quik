package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["_id"] = id
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractTokenMetadata(tokenString string) (bool, error) {
	token, err := jwt.Parse(extractToken(tokenString), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		_, ok := claims["_id"].(string)
		if !ok {
			return false, err
		}

		return true, nil
	}

	return false, err
}

func extractToken(bearerToken string) string {
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
