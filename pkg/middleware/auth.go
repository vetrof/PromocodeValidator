package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const userCtxKey ctxKey = "auth.user"

// User — структура, которую кладём в контекст
type User struct {
	ID    string
	Roles []string
}

// Claims — наши кастомные клеймы + RegisteredClaims
type Claims struct {
	Roles []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

type Auth struct {
	secret   []byte
	issuer   string
	audience string
}

func NewAuth(secret, issuer, audience string) *Auth {
	return &Auth{
		secret:   []byte(secret),
		issuer:   issuer,
		audience: audience,
	}
}

// Middleware — авторизация JWT.
// Если optional == false — токен обязателен (401 при ошибке).
// Если optional == true  — токен опционален (если есть и валиден, кладём User в ctx).
func (a *Auth) Middleware(optional bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tok, _ := bearerToken(r.Header.Get("Authorization"))
			if tok == "" {
				if optional {
					next.ServeHTTP(w, r) // токена нет — идём дальше без User
					return
				}
				http.Error(w, "missing or malformed Authorization header", http.StatusUnauthorized)
				return
			}

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tok, claims, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %T", t.Method)
				}
				return a.secret, nil
			})
			if err != nil || !token.Valid {
				if optional {
					next.ServeHTTP(w, r) // битый токен — идём дальше без User
					return
				}
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// issuer
			if a.issuer != "" && claims.Issuer != a.issuer {
				if !optional {
					http.Error(w, "invalid token issuer", http.StatusUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
				return
			}
			// audience
			if a.audience != "" {
				found := false
				for _, aud := range claims.Audience {
					if aud == a.audience {
						found = true
						break
					}
				}
				if !found {
					if !optional {
						http.Error(w, "invalid token audience", http.StatusUnauthorized)
						return
					}
					next.ServeHTTP(w, r)
					return
				}
			}
			// expiration
			if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
				if !optional {
					http.Error(w, "token expired", http.StatusUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
				return
			}

			uid := claims.Subject
			if uid == "" {
				if !optional {
					http.Error(w, "missing subject (user id)", http.StatusUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
				return
			}

			user := &User{ID: uid, Roles: claims.Roles}
			ctx := context.WithValue(r.Context(), userCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// UserFromContext — получить User в контроллере
func UserFromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userCtxKey).(*User)
	return u, ok
}

func bearerToken(h string) (string, error) {
	if h == "" {
		return "", errors.New("no header")
	}
	parts := strings.Fields(h)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("bad Authorization header")
	}
	return parts[1], nil
}

func (a *Auth) GenerateToken(userID string, roles []string, ttl time.Duration) (string, error) {
	claims := &Claims{
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    a.issuer,
			Audience:  []string{a.audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}
