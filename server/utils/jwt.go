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
var refreshkey=[]byte(os.Getenv("REFRESH_KEY"))

func GenerateToken(userID uint, role string) (string,string, error) {
	  // Token expires in 2 minutes
	 // Token expires in 2 minutes
  

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role, // Include role in token
		"exp":     time.Now().Add(29 * time.Minute).Unix(),
	}

	acesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	SignedaccessToken,err := acesstoken.SignedString(secretKey)

	if err !=nil{
		return "","",err
	}
 
	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role, // Include role in token
		"exp":     time.Now().Add(24 * time.Minute).Unix(),
	}

	refreshToken :=jwt.NewWithClaims(jwt.SigningMethodHS256,refreshClaims)
	SignedRefreshToken,err :=refreshToken.SignedString(refreshkey)


	if err !=nil{
		return "","",err
	}

	return SignedaccessToken,SignedRefreshToken,nil
}



// ValidateJWT verifies the given token
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
	})
	if err != nil || !token.Valid {
			return nil, errors.New("invalid tokens from validate token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
			return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

func ValidateRefreshToken(refreshtoken string)(jwt.MapClaims,error){
	token,err :=jwt.Parse(refreshtoken,func(token *jwt.Token)(interface{},error){
		return refreshkey,nil
	})

	if err !=nil{
		return nil,errors.New("invalid refresh token")
	}

	claims, ok :=token.Claims.(jwt.MapClaims)

	if !ok{
		return nil,errors.New("invalid refresh token claims")
	}
	return claims,nil
}
