package utils

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"uid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewAccessToken(secret, userID, role string, ttl time.Duration) (string, error) {
	claims := Claims{UserID: userID, Role: role, RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}

func NewRefreshToken(secret, userID string, ttl time.Duration) (string, error) {
	claims := Claims{UserID: userID, Role: "refresh", RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        fmt.Sprintf("%x", RandomBytes(8)),
	}}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}

func ParseToken(secret, token string) (*Claims, error) {
	tok, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token)(interface{},error){ return []byte(secret), nil })
	if err != nil { return nil, err }
	if c, ok := tok.Claims.(*Claims); ok && tok.Valid { return c, nil }
	return nil, jwt.ErrTokenInvalidClaims
}
