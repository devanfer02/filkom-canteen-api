package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	"github.com/devanfer02/filkom-canteen/internal/pkg/bcrypt"
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

	if flag.Flags.Seeder {
		log.Info(nil, "CONNECTION[NewPgsqlConn] Running seeders")
		runSeeder(dbx)
	}

	return dbx
}

func runSeeder(dbx *sqlx.DB) {
	var (
		err error 

		sqb squirrel.SelectBuilder
		iqb squirrel.InsertBuilder

		role  domain.Role
		admin = domain.Owner{
			Username: "adminfilkom",
			Password: "filkomevvah21",
			Fullname: "Admin FILKOM",
			WANumber: "+62 000000",
		}
	)

	admin.Password, _ = bcrypt.HashPassword(admin.Password)

	// this is not best practice but idgaf because is jus uni project
	roles := []map[string]interface{}{
		{"role_name": "Admin"},
		{"role_name": "Owner"},
	}

	_, err = dbx.NamedExec(`INSERT INTO roles (role_name) VALUES (:role_name)`, roles)

	sqb = squirrel.Select("*").From("roles")

	query, args, _ := sqb.PlaceholderFormat(squirrel.Dollar).ToSql()

	err = dbx.Get(&role, query)

	if err != nil {
		log.Warn(log.LogInfo{
			"err": err.Error(),
		}, "SEEDERS: failed to fetch role")
	}

	iqb = squirrel.
		Insert("admins").
		Columns("fullname", "username", "password", "wa_number", "role_id").
		Values(admin.Fullname, admin.Username, admin.Password, admin.WANumber, role.ID)

	query, args, _ = iqb.PlaceholderFormat(squirrel.Dollar).ToSql()

	_, err = dbx.Exec(query, args...)

	if err != nil {
		log.Warn(log.LogInfo{
			"err": err.Error(),
		}, "SEEDERS: failed to insert data")
	}
}
