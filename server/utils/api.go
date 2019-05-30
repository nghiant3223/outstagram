package utils

import (
	"os"
	"outstagram/server/models"
	"time"

	"outstagram/server/dtos"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AbortRequestWithError terminates request with error
func AbortRequestWithError(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "error", "message": message, "data": data})
}

// AbortRequestWithSuccess terminates request with success
func AbortRequestWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "success", "message": message, "data": data})
}

// SignToken returns a token string
func SignToken(user *models.User) (string, error) {
	claims := dtos.Token{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer:  os.Getenv("JWT_ISSUER"),
			Subject: user.Email, IssuedAt: time.Now().Unix(),
			ExpiresAt: int64(time.Now().Add(time.Duration(7 * 24 * time.Hour)).Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
