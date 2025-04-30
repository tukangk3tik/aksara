package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/tukangk3tik/aksara/api"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	setupLogger(config)
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

func setupLogger(config utils.Config) {
	logFile, _ := os.OpenFile(config.AppLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Configure encoder (JSON format)
	var core zapcore.Core
	encoderConfig := zap.NewProductionEncoderConfig()

	// Create core for file logging
	if config.AppEnv == "production" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Use JSON format
			zapcore.AddSync(multiWriter),          // Output to file
			zap.ErrorLevel,                        // Log level
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Use JSON format
			zapcore.AddSync(multiWriter),          // Output to file
			zap.DebugLevel,                        // Log level
		)
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	utils.GlobalLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	defer utils.GlobalLogger.Sync() // Flush logs before exiting
}
