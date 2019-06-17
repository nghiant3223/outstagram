package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"outstagram/server/utils"
	"strings"
)

func VerifyToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		utils.AbortRequestWithError(c, http.StatusUnauthorized, "No token provided", nil)
		return
	}

	if !strings.HasPrefix(tokenString, "Bearer") {
		utils.AbortRequestWithError(c, http.StatusBadRequest, "Unsupported token type", nil)
		return
	}

	tokenComponents := strings.Split(tokenString, " ")
	if len(tokenComponents) != 2 {
		utils.AbortRequestWithError(c, http.StatusBadRequest, "Invalid header's Authorization field", nil)
		return
	}

	token, err := jwt.Parse(tokenComponents[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		utils.AbortRequestWithError(c, http.StatusBadRequest, "Invalid token", err.Error())
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		utils.AbortRequestWithError(c, http.StatusBadRequest, "Invalid token", token)
		return
	}

	c.Set("userID", claims["userID"])
}
