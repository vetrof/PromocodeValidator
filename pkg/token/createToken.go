package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Generator struct {
	secret   []byte
	issuer   string
	audience string
}

func New(secret, issuer, audience string) *Generator {
	return &Generator{
		secret:   []byte(secret),
		issuer:   issuer,
		audience: audience,
	}
}

type Claims struct {
	Roles []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

func (g *Generator) Generate(userID string, roles []string, ttl time.Duration) (string, error) {
	claims := &Claims{
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    g.issuer,
			Audience:  []string{g.audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(g.secret)
}
