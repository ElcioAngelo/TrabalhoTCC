package model 

import 	"github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID int  `json:"user_id"`
	Email string `json: "email"`
	UserRole string `json: "user_role"`
	jwt.RegisteredClaims
}

