package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/devanfer02/filkom-canteen/internal/app/controller"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	"github.com/devanfer02/filkom-canteen/internal/middleware"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
)

type Server interface {
	MountMiddlewares()
	MountControllers()
	Start()
}

type httpServer struct {
	app *gin.Engine
	dbx *sqlx.DB
}

func NewHTTPServer(dbx *sqlx.DB) Server {
	app := gin.Default()

	return &httpServer{
		app: app,
		dbx: dbx,
	}
}

func (h *httpServer) MountMiddlewares() {
	h.app.Use(middleware.CORS())
}

func (h *httpServer) MountControllers() {
	v1 := h.app.Group("/api/v1")

	// repositories
	shopRepo := repository.NewShowRepository(h.dbx)

	// services
	shopSvc := service.NewShopService(shopRepo)

	// controllers
	controller.MountShopRoutes(v1, shopSvc)
}

func (h *httpServer) Start() {
	if env.AppEnv.AppPort[0] != ':' {
		env.AppEnv.AppPort = ":" + env.AppEnv.AppPort
	}

	if err := h.app.Run(env.AppEnv.AppPort); err != nil {
		
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "[HTTP SERVER][Start] failed to start server")
	}
}
