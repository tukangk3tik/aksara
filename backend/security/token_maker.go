package security

import "time"

type TokenMaker interface {
	GenerateToken(userId uint64, duration time.Duration) (string, *TokenPayload, error)
	VerifyToken(tokenString string) (*TokenPayload, error)
}