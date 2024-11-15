package main

import (
	database "github.com/devanfer02/filkom-canteen/internal/infra/database/pgsql"
	"github.com/devanfer02/filkom-canteen/internal/infra/server"
)

func main() {
	pgsqldb := database.NewPgsqlConn()
	httpSrv := server.NewHTTPServer(pgsqldb)

	httpSrv.MountMiddlewares()
	httpSrv.MountControllers()
	httpSrv.Start()
}