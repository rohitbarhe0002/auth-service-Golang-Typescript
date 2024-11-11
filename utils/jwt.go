package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT generates a JWT token with the provided email.
func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT token and returns the claims if valid.
func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}

// VerifyAndRefreshToken verifies a refresh token and generates a new access token if valid.
func VerifyAndRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// Check token validity
	if err != nil || !token.Valid {
		return "", errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}

	// Generate a new access token
	newClaims := jwt.MapClaims{
		"email": claims["email"],
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newAccessToken, err := newToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
