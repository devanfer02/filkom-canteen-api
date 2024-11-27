package service

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/google/uuid"
)

type IOrderService interface {
	FetchAllOrders() ([]domain.Order, error)
	FetchOrderByID(params *dto.OrderParams) (*domain.Order, error)
	CreateOrder(params *dto.OrderParams, req *dto.OrderRequest) error
	UpdateOrder(params *dto.OrderParams,req *dto.OrderRequest) error
	DeleteOrder(params *dto.OrderParams) error	
}

type orderServiceImpl struct {
	menuRepo repository.IOrderRepository
}

func NewOrderService(menuRepo repository.IOrderRepository) IOrderService {
	return &orderServiceImpl{menuRepo}
}

func(s *orderServiceImpl) FetchAllOrders() ([]domain.Order, error) {
	menus, err := s.menuRepo.FetchAll(&dto.OrderParams{})

	return menus, err 
}

func(s *orderServiceImpl) FetchOrderByID(params *dto.OrderParams) (*domain.Order, error) {
	if _, err := uuid.Parse(params.ID); err != nil {
		return nil, domain.ErrBadRequest
	}

	menu, err := s.menuRepo.FetchByID(params)

	return menu, err 
}

func(s *orderServiceImpl) CreateOrder(params *dto.OrderParams, req *dto.OrderRequest) error {
	err := s.menuRepo.InsertOrder(&domain.Order{
		UserID: req.UserID,
		MenuID: req.MenuID,
		
		Status: req.Status,
	})

	return err 
}

func(s *orderServiceImpl) UpdateOrder(params *dto.OrderParams,req *dto.OrderRequest) error {

	err := s.menuRepo.UpdateOrder(params, &domain.Order{
		Status: req.Status,
		PaymentMethod: req.PaymentMethod,
		PaymentProofLink: "",
	})

	return err 
}

func(s *orderServiceImpl) DeleteOrder(params *dto.OrderParams) error {
	err := s.menuRepo.DeleteOrder(params)

	return err 
}
