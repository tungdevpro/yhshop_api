package middleware

import (
	"coffee_api/configs"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtPayload struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func (jwt *JwtPayload) GetId() string {
	return jwt.Id
}

func GenToken(cfg *configs.Configuration, payload JwtPayload) (string, error) {
	secretKey := cfg.SecretKey
	claims := &JwtPayload{
		Id:   payload.Id,
		Role: payload.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3600).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return result, nil
}

func Validate(cfg *configs.Configuration, tokenString string) (string, error) {
	var claims JwtPayload

	tk, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if tk.Valid {
		// return tk.Claims.,
		fmt.Println("claims::: ", claims.Id)
	}

	return "", nil

	// jwt.ParseWithClaims(tokenString, &claims, )
}
