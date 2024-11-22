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

const MENU_TABLENAME = "menus"

type IMenuRepository interface {
	FetchAll(params *dto.MenuParams) ([]domain.Menu, error)
	FetchByID(params *dto.MenuParams) (*domain.Menu, error)
	InsertMenu(owner *domain.Menu) error
	UpdateMenu(params *dto.MenuParams, owner *domain.Menu) error
	DeleteMenu(params *dto.MenuParams) error
}

type menuRepositoryImpl struct {
	conn *sqlx.DB
}

func NewMehnuRepository(conn *sqlx.DB) IMenuRepository {
	return &menuRepositoryImpl{conn}
}

func (r *menuRepositoryImpl) FetchAll(params *dto.MenuParams) ([]domain.Menu, error) {
	var (
		qb     sq.SelectBuilder
		query  string
		args   []interface{}
		menus []domain.Menu = make([]domain.Menu, 0)
		err    error
	)

	qb = sq.Select("*").
		From(MENU_TABLENAME)

	query, args, err = qb.ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][FetchAll] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Select(&menus, query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][FetchAll] failed to fetch menus")

		return nil, err
	}

	return menus, nil
}

func (r *menuRepositoryImpl) FetchByID(params *dto.MenuParams) (*domain.Menu, error) {
	var (
		qb    sq.SelectBuilder
		query string
		args  []interface{}
		menu domain.Menu
		err   error
	)

	qb = sq.Select("*").
		From(MENU_TABLENAME).
		Where("menu_id = ?", params.ID).
		Limit(1)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][FetchByID] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Get(&menu, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}

		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][FetchByID] failed to fetch menu by id")

		return nil, err
	}

	return &menu, nil
}

func (r *menuRepositoryImpl) InsertMenu(menu *domain.Menu) error {
	var (
		qbi   sq.InsertBuilder
		query string
		err   error
		args  []any
	)

	qbi = sq.
		Insert(MENU_TABLENAME).
		Columns("menu_name", "shop_id", "menu_price", "menu_status", "menu_photo_link"). 
		Values(menu.Name, menu.ShopID, menu.Status, menu.PhotoLink)

	query, args, err = qbi.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][InsertMenu] failed to convert query builder to sql")
		return err
	}

	if _, err = r.conn.Exec(query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][InsertMenu] failed to execute sql statement")
		return err
	}

	return nil
}

func (r *menuRepositoryImpl) UpdateMenu(params *dto.MenuParams, menu *domain.Menu) error {
	var (
		qb    sq.UpdateBuilder
		query string
		err   error
		args  []any
	)

	qb = sq.
		Update(MENU_TABLENAME).
		Set("menu_name", menu.Name).
		Set("menu_price", menu.Status).
		Set("menu_photo_link", menu.PhotoLink).
		Set("updated_at", time.Now()). 
		Where("menu_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][UpdateMenu] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
			"query": query,
		}, "[MENU REPOSITORY][UpdateMenu] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *menuRepositoryImpl) DeleteMenu(params *dto.MenuParams) error {
	var (
		qb    sq.DeleteBuilder
		query string
		args  []interface{}
		err   error
	)

	qb = sq.
		Delete(MENU_TABLENAME).
		Where("menu_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][DeleteMenu] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][DeleteMenu] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}