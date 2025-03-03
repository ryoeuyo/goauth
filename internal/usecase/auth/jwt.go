package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ryoeuyo/goauth/internal/storage"
)

func NewToken(user *storage.User, ttl time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID.String()
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(ttl).Unix()

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, err
}
