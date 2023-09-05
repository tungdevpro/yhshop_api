package middleware

import (
	"coffee_api/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtPayload struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (jwt *JwtPayload) GetEmail() string {
	return jwt.Email
}

func GenToken(cfg *configs.Configuration, payload JwtPayload) (string, error) {
	secretKey := cfg.SecretKey
	claims := &JwtPayload{
		Email: payload.Email,
		Role:  payload.Role,
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
