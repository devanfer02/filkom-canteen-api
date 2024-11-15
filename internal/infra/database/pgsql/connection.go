package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
)

func NewPgsqlConn() *sqlx.DB {
	dbx, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=disable",
		env.AppEnv.DBUser,
		env.AppEnv.DBPass,
		env.AppEnv.DBHost,
		env.AppEnv.DBName,
	))

	if err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "[CONNECTION][NewPgsqlConn] failed to connect to database")
		
	}

	return dbx
}
