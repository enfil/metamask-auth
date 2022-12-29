package service

import (
	"fmt"
	"github.com/enfil/metamask-auth/domain/user"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Jwt struct {
	hmacSecret []byte
	issuer     string
	duration   time.Duration
	claims     jwt.RegisteredClaims
}

func NewJwt(hmacSecret string, issuer string, duration time.Duration) *Jwt {
	ans := Jwt{
		hmacSecret: []byte(hmacSecret),
		issuer:     issuer,
		duration:   duration,
	}
	return &ans
}

func (j *Jwt) Create(subject string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    j.issuer,
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(j.duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.hmacSecret)
}

func (j *Jwt) VerifyAndGetSubject(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.hmacSecret, nil
	})
	if err != nil {
		return "", user.ErrAuthError
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		j.claims = *claims
		return claims.Subject, nil
	}
	return "", user.ErrAuthError
}
