package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/devanfer02/filkom-canteen/internal/app/controller"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	"github.com/devanfer02/filkom-canteen/internal/middleware"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
	"github.com/devanfer02/filkom-canteen/internal/pkg/redis"
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
	// h.app.Use(middleware.APIKey()) // disabled for development
}

func (h *httpServer) MountControllers() {
	v1 := h.app.Group("/api/v1")
	redis := redis.NewRedisClient()

	url := ginSwagger.URL(env.AppEnv.AppUrl + `/swagger/doc.json`)

	// repositories
	shopRepo := repository.NewShowRepository(h.dbx)
	ownerRepo := repository.NewOwnerRepository(h.dbx)
	menuRepo := repository.NewMenuRepository(h.dbx)
	roleRepo := repository.NewRoleRepository(h.dbx)
	orderRepo := repository.NewOrderRepository(h.dbx)

	// middlewares
	mdlwr := middleware.NewMiddleware(redis, roleRepo)

	// services
	shopSvc := service.NewShopService(shopRepo)
	ownerSvc := service.NewOwnerService(ownerRepo)
	menuSvc := service.NewMenuService(menuRepo)
	orderSvc := service.NewOrderService(orderRepo)

	// controllers
	controller.MountShopRoutes(v1, shopSvc, mdlwr)
	controller.MountOwnerRoutes(v1, ownerSvc)
	controller.MountMenuRoutes(v1, menuSvc, mdlwr)
	controller.MountOrderRoutes(v1, orderSvc, mdlwr)

	h.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	h.app.GET("/hello", mdlwr.Authenticate(), mdlwr.AuthorizeAdmin("Owner"), func(ctx *gin.Context) {
		ctx.String(200, "Hello world")
	})
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
