package security

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	"github.com/tukangk3tik/aksara/utils"
)

func TestJwtTokenMaker(t *testing.T) {
	maker, err := NewJwtTokenMaker(utils.RandomString(32))
	require.NoError(t, err)
	
	id := utils.GenerateSnowflakeID()
	payload, err := NewPayload(id, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	token, payload, err := maker.GenerateToken(id, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, id, payload.UserId)
	require.WithinDuration(t, payload.RegisteredClaims.ExpiresAt.Time, time.Now(), time.Minute)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJwtTokenMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, _, err := maker.GenerateToken(utils.GenerateSnowflakeID(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	fmt.Println(err)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(utils.GenerateSnowflakeID(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJwtTokenMaker(utils.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}

func TestMinSecretKeySize(t *testing.T) {
	_, err := NewJwtTokenMaker(utils.RandomString(10))
	require.Error(t, err)
	require.EqualError(t, err, fmt.Sprintf("invalid key size: must be at least %d characters", minSecretKeySize))
}