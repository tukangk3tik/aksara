package main

import (
	"database/sql"
	"log"

	"github.com/tukangk3tik/aksara/api"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	utils.SetupLogger(config)
	runGatewayServer(config, store)
}

func runGatewayServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		utils.GlobalLogger.Fatal("cannot create server:", zap.Error(err))
	}

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		utils.GlobalLogger.Fatal("cannot start server:", zap.Error(err))
	}
}

