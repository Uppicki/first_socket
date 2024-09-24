package domain

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	Login string `json:"login"`
	Type  string `json:"type"`
	jwt.StandardClaims
}
