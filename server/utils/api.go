package utils

import (
	"os"
	"outstagram/server/models"
	"time"

	"outstagram/server/dtos"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AbortRequestWithError abort request with error, request stops at middleware in which this function is called
func AbortRequestWithError(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "error", "message": message, "data": data})
}

// AbortRequestWithSuccess abort request with success, request stops at middleware in which this function is called
func AbortRequestWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "success", "message": message, "data": data})
}

// ResponseWithError responses request with an error
func ResponseWithError(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{"status": "error", "message": message, "data": data})
}

// ResponseWithSuccess responses request with a success
func ResponseWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{"status": "success", "message": message, "data": data})
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
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
