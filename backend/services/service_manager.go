package services

import (
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/security"
	"github.com/tukangk3tik/aksara/utils"

	_ "github.com/go-sql-driver/mysql"
)

type ServiceManager struct {
	config     utils.Config
	store      db.Store
	tokenMaker security.TokenMaker
}

func NewServiceManager(config utils.Config, store db.Store, tokenMaker security.TokenMaker) *ServiceManager {
	return &ServiceManager{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
}
