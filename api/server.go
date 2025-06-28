package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/security"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
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

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	/*
		if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		}
	*/

	if server.config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	server.setupRouter(tokenMaker)
	return server, nil
}

func (server *Server) setupRouter(tokenMaker security.TokenMaker) {
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

	router.POST("/auth/login", server.loginUser)

	// location router
	locGroup := router.Group("/loc")
	locGroup.Use(security.AuthorizeJwt(tokenMaker))
	locGroup.GET("/provinces", server.fetchProvinces)
	locGroup.GET("/regencies", server.fetchRegencyByProvince)
	locGroup.GET("/districts", server.fetchDistrictByRegency)

	// office router
	officeGroup := router.Group("/offices")
	officeGroup.Use(security.AuthorizeJwt(tokenMaker))
	officeGroup.GET("", server.getOffices)
	officeGroup.POST("", server.createOffice)
	officeGroup.PUT("/:id", server.updateOffice)
	officeGroup.DELETE("/:id", server.deleteOffice)
	officeGroup.GET("/select-option", server.fetchOfficesSelectOption)

	// school router
	schoolGroup := router.Group("/schools")
	schoolGroup.Use(security.AuthorizeJwt(tokenMaker))
	schoolGroup.GET("", server.getSchools)
	schoolGroup.GET("/:id", server.getSchoolById)
	schoolGroup.POST("", server.createSchool)
	schoolGroup.PUT("/:id", server.updateSchool)
	schoolGroup.DELETE("/:id", server.deleteSchool)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
