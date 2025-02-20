package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/security"
	"github.com/tukangk3tik/aksara/services"
	"github.com/tukangk3tik/aksara/utils"
)

type Server struct {
	config utils.Config
	sm     *services.ServiceManager
	router *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := security.NewJwtTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	fmt.Println(config)
	sm := services.NewServiceManager(config, store, tokenMaker)
	server := &Server{
		config: config,
		sm:     sm,
	}
	/*
		if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		}
	*/

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, SuccessResponse{
			StatusCode: 200,
			Message:    "Aksara API - v1",
		})
	})

	router.POST("/users/login", server.loginUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
