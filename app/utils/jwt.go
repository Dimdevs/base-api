package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var jwtKey = []byte("secret")

func CreateToken(userID uuid.UUID, role string) (string, error) {
	userIDString := userID.String()

	claims := &jwt.StandardClaims{
		Subject:   userIDString,
		Issuer:    role,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
