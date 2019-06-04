package jwtdtos

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID uint `json:"userID"`
	jwt.StandardClaims
}
