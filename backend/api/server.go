package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/security"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	config     utils.Config
	router     *gin.Engine
	store      db.Store
	logger     *zap.Logger
	tokenMaker security.TokenMaker
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := security.NewJwtTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	fmt.Println(config)
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	/*
		if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		}
	*/

	server.setupLogger()
	server.setupRouter(tokenMaker)
	return server, nil
}

func (server *Server) setupLogger() {
	logFile, _ := os.OpenFile(server.config.AppLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileSyncer := zapcore.AddSync(logFile)

	// Configure encoder (JSON format)
	var encoderConfig zapcore.EncoderConfig
	var core zapcore.Core

	// Create core for file logging
	if server.config.AppEnv == "production" {
		encoderConfig = zap.NewProductionEncoderConfig()
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Use JSON format
			fileSyncer,                            // Output to file
			zap.InfoLevel,                         // Log level
		)
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Use JSON format
			fileSyncer,                            // Output to file
			zap.DebugLevel,                        // Log level
		)
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(core)
	defer logger.Sync() // Flush logs before exiting

	server.logger = logger
}

func (server *Server) setupRouter(tokenMaker security.TokenMaker) {
	if server.config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.Use(security.TraceMiddleware(), security.CorsMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response.SuccessResponse{
			Data: "Aksara API - v1",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.BuildErrorResponse("NOT_FOUND", utils.ErrorCodeMap["NOT_FOUND"], nil))
	})

	router.POST("/users/login", server.loginUser)

	officeGroup := router.Group("/offices")
	officeGroup.Use(security.AuthorizeJwt(tokenMaker))
	officeGroup.GET("/", server.getOffices)
	officeGroup.POST("/", server.createOffice)
	officeGroup.PUT("/:id", server.updateOffice)
	officeGroup.DELETE("/:id", server.deleteOffice)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
