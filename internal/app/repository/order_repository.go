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

const ORDER_TABLENAME = "orders"

type IOrderRepository interface {
	FetchAll(params *dto.OrderParams) ([]domain.Order, error)
	FetchByID(params *dto.OrderParams) (*domain.Order, error)
	InsertOrder(owner *domain.Order) error
	UpdateOrder(params *dto.OrderParams, owner *domain.Order) error
	DeleteOrder(params *dto.OrderParams) error
}

type orderRepositoryImpl struct {
	conn *sqlx.DB
}

func NewOrderRepository(conn *sqlx.DB) IOrderRepository {
	return &orderRepositoryImpl{conn}
}

func (r *orderRepositoryImpl) FetchAll(params *dto.OrderParams) ([]domain.Order, error) {
	var (
		qb     sq.SelectBuilder
		query  string
		args   []interface{}
		orders []domain.Order = make([]domain.Order, 0)
		err    error
	)

	qb = sq.Select("*").
		From(ORDER_TABLENAME)

	query, args, err = qb.ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][FetchAll] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Select(&orders, query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][FetchAll] failed to fetch orders")

		return nil, err
	}

	return orders, nil
}

func (r *orderRepositoryImpl) FetchByID(params *dto.OrderParams) (*domain.Order, error) {
	var (
		qb    sq.SelectBuilder
		query string
		args  []interface{}
		order domain.Order
		err   error
	)

	qb = sq.Select("*").
		From(ORDER_TABLENAME).
		Where("order_id = ?", params.ID).
		Limit(1)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][FetchByID] failed to convert query builder to sql")
		return nil, err
	}

	if err = r.conn.Get(&order, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}

		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][FetchByID] failed to fetch order by id")

		return nil, err
	}

	return &order, nil
}

func (r *orderRepositoryImpl) InsertOrder(order *domain.Order) error {
	var (
		qbi   sq.InsertBuilder
		query string
		err   error
		args  []any
	)

	qbi = sq.
		Insert(ORDER_TABLENAME).
		Columns("user_id", "menu_id", "payment_method", "status", "payment_proof_link"). 
		Values(order.UserID, order.MenuID, order.PaymentMethod, order.Status, order.PaymentProofLink)

	query, args, err = qbi.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][InsertOrder] failed to convert query builder to sql")
		return err
	}

	if _, err = r.conn.Exec(query, args...); err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][InsertOrder] failed to execute sql statement")
		return err
	}

	return nil
}

func (r *orderRepositoryImpl) UpdateOrder(params *dto.OrderParams, order *domain.Order) error {
	var (
		qb    sq.UpdateBuilder
		query string
		err   error
		args  []any
	)

	qb = sq.
		Update(ORDER_TABLENAME).
		Set("status", order.Status). 
		Set("payment_method", order.PaymentMethod). 
		Set("payment_proof_link", order.PaymentProofLink). 
		Set("updated_at", time.Now()). 
		Where("order_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][UpdateOrder] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][UpdateOrder] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *orderRepositoryImpl) DeleteOrder(params *dto.OrderParams) error {
	var (
		qb    sq.DeleteBuilder
		query string
		args  []interface{}
		err   error
	)

	qb = sq.
		Delete(ORDER_TABLENAME).
		Where("order_id = ?", params.ID)

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][DeleteOrder] failed to convert query builder to sql")
		return err
	}

	res, err := r.conn.Exec(query, args...)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[ORDER REPOSITORY][DeleteOrder] failed to execute sql statement")
		return err
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}
