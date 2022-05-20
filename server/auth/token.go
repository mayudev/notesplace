package auth

import "github.com/golang-jwt/jwt"

type Claims struct {
	NotebookID string `json:"notebook_id"`
	jwt.StandardClaims
}
