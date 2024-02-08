package security

import (
	"fmt"
	"time"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateAccesToken(email string) (string, error) {
	cfg := config.LoadEnv()
	claims := JwtCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(cfg.JWTSecret))

	if err != nil {
		fmt.Println(err)
		return t, err
	}

	return t, err
}
