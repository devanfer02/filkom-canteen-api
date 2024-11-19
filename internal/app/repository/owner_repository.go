package repository

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
)

const OWNER_TABLENAME = "admins"

type IOwnerRepository interface {
	FetchAll(params *dto.OwnerParams) ([]domain.Owner, error)
	FetchByID(params *dto.OwnerParams) (*domain.Owner, error)
	InsertOwner(owner *domain.Owner) error
	UpdateOwner(params *dto.OwnerParams, owner *domain.Owner) error
	DeleteOwner(params *dto.OwnerParams) error
}

type ownerRepositoryImpl struct {
	conn *sqlx.DB
}

func NewOwnerRepository(conn *sqlx.DB) IOwnerRepository {
	return &ownerRepositoryImpl{conn}
}

func (r *ownerRepositoryImpl) FetchAll(params *dto.OwnerParams) ([]domain.Owner, error) {
	var (
		qb    sq.SelectBuilder
		query string
		args  []interface{}
		owners []domain.Owner = make([]domain.Owner, 0)
		err   error
	)

	qb = sq.Select("admin_id", "fullname", "wa_number", "username", "password", "created_at", "updated_at"). 
		From(OWNER_TABLENAME). 
		Join("roles ON roles.role_id = admins.role_id").
		Where("roles.role_name = ?", "Owner")

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][FetchAll] failed to convert query builder to sql")
		return nil, err	
	}

	if err = r.conn.Select(&owners, query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][FetchAll] failed to fetch owners")

		return nil, err 
	} 

	return owners, nil 
}

func (r *ownerRepositoryImpl) FetchByID(params *dto.OwnerParams) (*domain.Owner, error) {
	var (
		qb    sq.SelectBuilder
		query string
		args  []interface{}
		owner domain.Owner
		err   error
	)

	qb = sq.Select("admin_id", "fullname", "wa_number", "username", "password", "created_at", "updated_at"). 
		From(OWNER_TABLENAME).
		Join("roles ON roles.role_id = admins.role_id").
		Where("admin_id = ? AND roles.role_name = ?", params.ID, "Owner"). 
		Limit(1)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][FetchByID] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Get(&owner, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}

		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][FetchByID] failed to fetch owner by id")

		return nil, err 
	}

	return &owner, nil 
}

func (r *ownerRepositoryImpl) InsertOwner(owner *domain.Owner) error {
	var (
		qbi    sq.InsertBuilder
		qbs    sq.SelectBuilder
		query string
		err   error
		args  []any
		role  domain.Role
	)

	qbs = sq.
		Select("*").
		From("roles").
		Where("role_name = ?", "Owner"). 
		Limit(1)

	query, args, _ = qbs.PlaceholderFormat(sq.Dollar).ToSql()

	err = r.conn.Get(&role, query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][InsertOwner] failed to fetch role")
		return err
	}

	log.Info(log.LogInfo{
		"owner": owner,
	}, "LOG OWNER")

	qbi = sq.
		Insert(OWNER_TABLENAME).
		Columns("fullname", "wa_number", "username", "password", "role_id").
		Values(owner.Fullname, owner.WANumber, owner.Username, owner.Password, role.ID)

	query, args, err = qbi.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][InsertOwner] failed to convert query builder to sql")
		return err
	}

	if _, err = r.conn.Exec(query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][InsertOwner] failed to execute sql statement")
		return err
	}

	return nil
}

func (r *ownerRepositoryImpl) UpdateOwner(params *dto.OwnerParams, owner *domain.Owner) error {
	var (
		qb    sq.UpdateBuilder
		query string
		err   error
		args  []any
	)

	qb = sq.
		Update(OWNER_TABLENAME).
		Set("fullname", owner.Fullname).
		Set("wa_number", owner.WANumber).
		Set("username", owner.Username).
		Set("updated_at", time.Now())

	if owner.Password != "" {
		qb = qb.Set("password", owner.Password)
	}
		
	qb = qb.Where("admin_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][UpdateOwner] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
			"query": query,
		}, "[OWNER REPOSITORY][UpdateOwner] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *ownerRepositoryImpl) DeleteOwner(params *dto.OwnerParams) error {
	var (
		qb sq.DeleteBuilder
		query string 
		args []interface{}
		err error 
	)

	qb = sq.
		Delete(OWNER_TABLENAME).
		Where("admin_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()


	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][DeleteOwner] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[OWNER REPOSITORY][DeleteOwner] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}
