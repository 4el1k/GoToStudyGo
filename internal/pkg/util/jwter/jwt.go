package jwter

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTer interface {
	EncodeToken(string) (string, error)
	DecodeToken(string) (string, error)
}

type authClaims struct {
	principal string
	jwt.RegisteredClaims
}

func EncodeToken(username string) (string, time.Time, error) {
	exp := time.Now()
	authClaims := &authClaims{
		principal: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			Issuer:    "test_issuer",
			ExpiresAt: jwt.NewNumericDate(exp.Add(time.Minute * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	t, err := token.SignedString([]byte("some_secret"))
	if err != nil {
		return "", time.Time{}, err
	}
	return t, exp, nil
}

func DecodeToken(jwtToken string) (string, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &authClaims{}, getKeyFunc())
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*authClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	return claims.Subject, nil
}

func getKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("some_secret"), nil
	}
}
