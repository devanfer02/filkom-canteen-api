package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
	"github.com/jmoiron/sqlx"
)

const ROLE_TABLENAME = "roles"

type IRoleRepository interface {
	FetchOne(name string) (*domain.Role, error)
}

type roleRepositoryImpl struct {
	conn *sqlx.DB
}

func NewRoleRepository(conn *sqlx.DB) IRoleRepository {
	return &roleRepositoryImpl{conn}
}

func (r *roleRepositoryImpl) FetchOne(name string) (*domain.Role, error) {
	var (
		qb sq.SelectBuilder
		query string 
		args []any
		role domain.Role
		err error 
	)

	qb = sq.Select("*").From(ROLE_TABLENAME).Where("role_name = ?", name).Limit(1)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ROLE REPOSITORY][FetchOne] failed to fetch role")
		return nil, err 
	}

	if err = r.conn.Get(&role, query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ROLE REPOSITORY][FetchOne] failed to fetch role")
		return nil, err 
	}

	return &role, nil 
}