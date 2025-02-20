package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tukangk3tik/aksara/security"
	"github.com/tukangk3tik/aksara/utils"
)

type LoginResult struct {
	AccessToken  string
	RefreshToken string
}

func (sm *ServiceManager) LoginUser(ctx context.Context, email string, password string) (result LoginResult, err error) {
	user, err := sm.store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("user tidak terdaftar")
		}
		return result, err
	}

	err = utils.CheckPassword(password, user.Password)
	if err != nil {
		return result, errors.New("password salah")
	}

	tokenMaker, err := security.NewJwtTokenMaker(sm.config.TokenSymmetricKey)
	if err != nil {
		return result, err
	}

	accessToken, _, err := tokenMaker.GenerateToken(user.ID, sm.config.AccessTokenDuration)
	if err != nil {
		return result, err
	}
	
	refreshToken, _, err := tokenMaker.GenerateToken(user.ID, sm.config.RefreshTokenDuration)
	if err != nil {
		return result, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken
	
	return
}
