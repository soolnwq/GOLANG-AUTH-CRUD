package utils

import (
	"go-crud/errs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateAccessToken(userId int, exp time.Time) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     exp.Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.access_secret")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func VerifyAccessToken(tokenString string) (*jwt.MapClaims, error) {
	unauth := func(msg string) error {
		return errs.NewUnautherizedError(msg)
	}

	if tokenString == "" {
		return nil, unauth("access token is missing from request")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.access_secret")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil || !token.Valid {
		return nil, unauth("access token is invalid or malformed")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, unauth("access token has invalid claims format")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, unauth("access token missing 'exp' (expiration) claim")
	}

	if time.Now().Unix() > int64(exp) {
		return nil, unauth("access token has expired")
	}

	return &claims, nil
}
