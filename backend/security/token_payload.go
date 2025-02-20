package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token signature is invalid: token is unverifiable: 'none' signature type is not allowed")
	ErrExpiredToken = errors.New("token has invalid claims: token is expired")
)

type TokenPayload struct {
	ID     uuid.UUID `json:"id"`
	UserId uint64    `json:"user_id"`
	jwt.RegisteredClaims
}

func NewPayload(userId uint64, duration time.Duration) (*TokenPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &TokenPayload{
		tokenID,
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "aksara",
			Subject:   "aksara",
			ID:        "1",
			Audience:  []string{"aksara"},
		},
	}
	return payload, nil
}

func (payload *TokenPayload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
