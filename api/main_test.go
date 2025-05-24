package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/security"
	"github.com/tukangk3tik/aksara/utils"
)

const provinceID = 53
const regencyID = 5371
const districtID = 5371010

type mockSQLResult struct{}

func (m mockSQLResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (m mockSQLResult) RowsAffected() (int64, error) {
	return 1, nil
}

type mockNotFoundSQLResult struct{}

func (m mockNotFoundSQLResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (m mockNotFoundSQLResult) RowsAffected() (int64, error) {
	return 0, nil
}

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:    utils.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
		AppEnv:               "test",
		LogLevel:             "info",
	}

	tokenMaker, err := security.NewJwtTokenMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter(tokenMaker)
	utils.SetupLogger(config)	
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}