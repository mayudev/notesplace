package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	NotebookID string `json:"notebook_id"`
	jwt.StandardClaims
}

type issuer struct {
	SigningKey []byte
}

// NewIssuer returns a new configured JWT issuer
func NewIssuer(key string) *issuer {
	issuer := issuer{}
	issuer.SigningKey = []byte(key)

	return &issuer
}

// Issue issues a new JWT
func (issuer *issuer) Issue(id string) (string, error) {
	claims := Claims{
		NotebookID: id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().UTC().Unix(),
			ExpiresAt: time.Now().UTC().Add(3 * 24 * time.Hour).Unix(),
			Issuer:    "notesplace",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signed, err := token.SignedString(issuer.SigningKey)

	if err != nil {
		return "", err
	}

	return signed, nil
}

// Validate tries to validate a JWT, returning Claims or an error
func (issuer *issuer) Validate(token string) (*Claims, error) {
	parsed, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return issuer.SigningKey, nil
		},
	)

	if err != nil {
		return &Claims{}, err
	}

	claims, ok := parsed.Claims.(*Claims)
	if !ok {
		return &Claims{}, fmt.Errorf("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return &Claims{}, fmt.Errorf("token has expired")
	}

	return claims, nil
}

// ValidateNotebook validates if a token grants authorized access to a notebook
func (issuer *issuer) ValidateNotebook(token string, notebookID string) bool {
	validated, err := issuer.Validate(token)
	if err != nil {
		return false
	}

	if validated.NotebookID == notebookID {
		return true
	}

	return false
}
