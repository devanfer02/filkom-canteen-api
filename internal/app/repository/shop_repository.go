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

const TABLE_NAME = "shops"

type IShopRepository interface {
	FetchAllShops() ([]domain.Shop, error)
	FetchShopByID(params *dto.ShopParams) (*domain.Shop, error)
	InsertShop(shop *domain.Shop) error
	UpdateShop(params *dto.ShopParams, shop *domain.Shop) error
	DeleteShop(params *dto.ShopParams) error
}

type shopRepositoryImpl struct {
	conn *sqlx.DB
}

func NewShowRepository(conn *sqlx.DB) IShopRepository {
	return &shopRepositoryImpl{conn: conn}
}

func (r *shopRepositoryImpl) FetchAllShops() ([]domain.Shop, error) {
	var (
		qb    sq.SelectBuilder
		query string
		err   error
		shops []domain.Shop = make([]domain.Shop, 0)
	)

	qb = sq.Select("*").From(TABLE_NAME)

	query, _, err = qb.ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][FetchAllShops] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Select(&shops, query); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][FetchAllShops] failed to fetch shops")
		return nil, err
	}

	return shops, nil
}

func (r *shopRepositoryImpl) FetchShopByID(params *dto.ShopParams) (*domain.Shop, error) {
	var (
		qb    sq.SelectBuilder
		query string
		err   error
		shop  domain.Shop
		args  []interface{}
	)

	qb = sq.Select("*").From(TABLE_NAME).Where("shop_id = ?", params.ID).Limit(1)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][FetchShopByID] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Get(&shop, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}

		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][FetchShopByID] failed to fetch shop")
		return nil, err
	}

	return &shop, nil
}

func (r *shopRepositoryImpl) InsertShop(shop *domain.Shop) error {
	var (
		qb    sq.InsertBuilder
		query string
		err   error
		args  []any
	)

	qb = sq.
		Insert(TABLE_NAME).
		Columns("shop_name", "shop_description", "shop_photo_link").
		Values(shop.Name, shop.Description, shop.PhotoLink)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][InsertShop] failed to convert query builder to sql")
		return err
	}

	if _, err = r.conn.Exec(query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][InsertShop] failed to execute sql statement")
		return err
	}

	return nil

}

func (r *shopRepositoryImpl) UpdateShop(params *dto.ShopParams, shop *domain.Shop) error {
	var (
		qb    sq.UpdateBuilder
		query string
		err   error
		args  []any
	)

	qb = sq.
		Update(TABLE_NAME).
		Set("shop_name", shop.Name).
		Set("shop_description", shop.Description).
		Set("shop_photo_link", shop.PhotoLink).
		Set("updated_at", time.Now()).
		Where("shop_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][UpdateShop] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
			"query": query,
		}, "[SHOP REPOSITORY][UpdateShop] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *shopRepositoryImpl) DeleteShop(params *dto.ShopParams) error {
	var (
		qb    sq.DeleteBuilder
		query string
		err   error
		args  []any
	)

	qb = sq.
		Delete(TABLE_NAME).
		Where("shop_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][DeleteShop] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[SHOP REPOSITORY][DeleteShop] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}
