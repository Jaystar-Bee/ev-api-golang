package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "MadebyJaystar"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 6).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (jwt.MapClaims, error) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("failed to parse token")
	}

	isValidToken := parsedToken.Valid

	if !isValidToken {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	// expireTime := claims["exp"].(int64)
	// if expireTime > time.Now().Unix() {
	// 	return nil, errors.New("jwt.expired")
	// }

	return claims, nil
}
