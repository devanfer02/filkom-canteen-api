package repository

import (
	"database/sql"
	"strings"
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

func NewMenuRepository(conn *sqlx.DB) IMenuRepository {
	return &menuRepositoryImpl{conn}
}

func (r *menuRepositoryImpl) FetchAll(params *dto.MenuParams) ([]domain.Menu, error) {
	var (
		qb    sq.SelectBuilder
		query string
		args  []interface{}
		menus []domain.Menu = make([]domain.Menu, 0)
		err   error
	)

	qb = sq.Select(
		"menu_id",
		"menu_name",
		"menus.shop_id AS menu_shop_id",
		"menu_price",
		"menu_status",
		"menu_photo_link",
		"menus.created_at AS created_at",
		"menus.updated_at AS updated_at",
	).From(MENU_TABLENAME)

	if params.ShopID != "" {
		qb = qb.
			Join("shops ON shops.shop_id = menus.shop_id").
			Where("shops.shop_id = ?", params.ShopID)
	}

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][FetchAll] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Select(&menus, query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
			"query": query,
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
		menu  domain.Menu
		err   error
	)

	qb = sq.Select(
		"menu_id",
		"menu_name",
		"menus.shop_id AS menu_shop_id",
		"menu_price",
		"menu_status",
		"menu_photo_link",
		"created_at",
		"updated_at",
	).From(MENU_TABLENAME).
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
		Values(menu.Name, menu.ShopID, menu.Price, menu.Status, menu.PhotoLink)

	query, args, err = qbi.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[MENU REPOSITORY][InsertMenu] failed to convert query builder to sql")
		return err
	}

	if _, err = r.conn.Exec(query, args...); err != nil {

		if strings.Contains(err.Error(), "violates") {
			return domain.ErrBadRequest
		}

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
		Set("menu_price", menu.Price).
		Set("menu_photo_link", menu.PhotoLink).
		Set("menu_status", menu.Status).
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
