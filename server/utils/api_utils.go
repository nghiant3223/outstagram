package utils

import (
	"os"

	"outstagram/server/dtos"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AbortRequestWithError terminates request with error
func AbortRequestWithError(c *gin.Context, code int, message interface{}, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "error", "message": message, "data": data})
}

// AbortRequestWithSuccess terminates request with success
func AbortRequestWithSuccess(c *gin.Context, code int, message interface{}, data interface{}) {
	c.JSON(code, gin.H{"status": "success", "message": message, "data": data})
}

// SignToken returns a token string
func SignToken(userID uint) (string, error) {
	claims := dtos.Token{UserID: userID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
