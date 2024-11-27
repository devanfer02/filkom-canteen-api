package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	"github.com/devanfer02/filkom-canteen/internal/pkg/flag"
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

	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})

	if err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "CONNECTION[NewPgsqlConn]")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		env.AppEnv.DBName, driver,
	)

	if err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "CONNECTION[NewPgsqlConn]")
	}

	if flag.Flags.Fresh {
		log.Info(nil, "CONNECTION[NewPgsqlConn] Dropping all tables")
		m.Down()
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "CONNECTION[NewPgsqlConn]")
	}

	return dbx
}
