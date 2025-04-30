package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JwtTokenMaker struct {
	secretKey string
}

func NewJwtTokenMaker(secretKey string) (TokenMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JwtTokenMaker{secretKey: secretKey}, nil
}

func (maker *JwtTokenMaker) GenerateToken(userId uint64, duration time.Duration) (string, *TokenPayload, error) {
	payload, err := NewPayload(userId, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := jwtToken.SignedString([]byte(maker.secretKey))
	return signedToken, payload, err
}

func (maker *JwtTokenMaker) VerifyToken(tokenString string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(maker.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenPayload)
	if !ok {
		return nil, errors.New("unknown claims type, cannot proceed")
	}

	return claims, nil
}
