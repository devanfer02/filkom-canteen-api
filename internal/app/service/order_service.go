package service

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	enc "github.com/devanfer02/filkom-canteen/internal/pkg/encoder"
	"github.com/google/uuid"
)

type IOrderService interface {
	FetchAllOrders(params *dto.OrderParams) ([]domain.Order, error)
	FetchOrderByID(params *dto.OrderParams) (*domain.Order, error)
	CreateOrder(params *dto.OrderParams, req *dto.OrderRequest) error
	UpdateOrder(params *dto.OrderParams, req *dto.OrderRequest) error
	DeleteOrder(params *dto.OrderParams) error
}

type orderServiceImpl struct {
	orderRepo repository.IOrderRepository
}

func NewOrderService(orderRepo repository.IOrderRepository) IOrderService {
	return &orderServiceImpl{orderRepo}
}

func (s *orderServiceImpl) FetchAllOrders(params *dto.OrderParams) ([]domain.Order, error) {
	orders, err := s.orderRepo.FetchAll(params)

	if err != nil {
		return nil, err 
	}

	for idx, order := range orders {
		orders[idx].ID = enc.Encode(order.ID)
		orders[idx].MenuID = enc.Encode(order.MenuID)
	}

	return orders, err
}

func (s *orderServiceImpl) FetchOrderByID(params *dto.OrderParams) (*domain.Order, error) {
	decoded, err := enc.Decode(params.ID)

	if err != nil {
		return nil, domain.ErrBadRequest
	}

	params.ID = decoded

	if _, err := uuid.Parse(params.ID); err != nil {
		return nil, domain.ErrBadRequest
	}
	
	order, err := s.orderRepo.FetchByID(params)

	if err != nil {
		return nil, err 
	}

	order.ID = enc.Encode(order.ID) 
	order.MenuID = enc.Encode(order.MenuID)

	return order, err
}

func (s *orderServiceImpl) CreateOrder(params *dto.OrderParams, req *dto.OrderRequest) error {
	decodedMenuID, err := enc.Decode(req.MenuID)

	if err != nil {
		return domain.ErrBadRequest
	}

	err = s.orderRepo.InsertOrder(&domain.Order{
		UserID: params.UserID,
		MenuID: decodedMenuID,
		PaymentMethod: req.PaymentMethod,
		Status: "Waiting",
	})

	return err
}

func (s *orderServiceImpl) UpdateOrder(params *dto.OrderParams, req *dto.OrderRequest) error {

	err := s.orderRepo.UpdateOrder(params, &domain.Order{
		Status:           req.Status,
		PaymentMethod:    req.PaymentMethod,
		PaymentProofLink: "",
	})

	return err
}

func (s *orderServiceImpl) DeleteOrder(params *dto.OrderParams) error {
	err := s.orderRepo.DeleteOrder(params)

	return err
}
