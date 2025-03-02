package utils

import (
	"errors"
	// "fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Custom claims struct
type Claims struct {
	Username    string   `json:"username"`

	jwt.RegisteredClaims
}

// GenerateJWT generates a new token for a given username, role, and permissions



var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(userID uint, role string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Minute) // Token expires in 2 minutes

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role, // Include role in token
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}



// ValidateJWT verifies the given token
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
	})
	if err != nil || !token.Valid {
			return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
			return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
