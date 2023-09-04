package middleware

import (
	"coffee_api/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (jwt *JwtPayload) GetUserId() int {
	return jwt.UserId
}

func GenToken(cfg *configs.Configuration, payload JwtPayload) (string, error) {
	secretKey := cfg.SecretKey
	claims := &JwtPayload{
		UserId: payload.UserId,
		Role:   payload.Role,
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
