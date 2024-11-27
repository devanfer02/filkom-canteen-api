package main

import (
	_ "github.com/devanfer02/filkom-canteen/docs"

	database "github.com/devanfer02/filkom-canteen/internal/infra/database/pgsql"
	"github.com/devanfer02/filkom-canteen/internal/infra/server"
)

//	@title						FILKOM Canteen API
//	@version					1.0
//	@description				This is FILKOM Canteen API Documentation
//	@host						localhost:5700
//	@@host						filkom-api.dvnnfrr.my.id
//	@@schemes					https
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						x-api-key
//	@description				API Key for accessing all endpoints. Type: Key TOKEN
func main() {
	pgsqldb := database.NewPgsqlConn()
	httpSrv := server.NewHTTPServer(pgsqldb)

	httpSrv.MountMiddlewares()
	httpSrv.MountControllers()
	httpSrv.Start()
}