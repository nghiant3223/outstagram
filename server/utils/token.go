package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"outstagram/server/dtos/jwtdtos"
	"outstagram/server/models"
	"time"
)

// SignToken returns a token string
func SignToken(user *models.User) (string, error) {
	claims := jwtdtos.Token{
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

func RetrieveUserID(c *gin.Context) (uint, bool) {
	userID, ok := c.Get("userID")
	if !ok {
		return 0, false
	}

	return uint(userID.(float64)), true
}
