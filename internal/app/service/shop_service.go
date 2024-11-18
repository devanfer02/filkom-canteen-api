package service

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/google/uuid"
)

type IShopService interface {
	FetchAllShops() ([]domain.Shop, error)
	FetchShopByID(params *dto.ShopParams) (*domain.Shop, error)
	CreateShop(req *dto.ShopRequest) error
	UpdateShop(req *dto.ShopRequest) error
	DeleteShop(params *dto.ShopParams) error
}

type shopServiceImpl struct {
	shopRepo repository.IShopRepository
}

func NewShopService(shopRepo repository.IShopRepository) IShopService {
	return &shopServiceImpl{shopRepo: shopRepo}
}

func (s *shopServiceImpl) FetchAllShops() ([]domain.Shop, error) {
	shops, err := s.shopRepo.FetchAllShops()

	return shops, err
}

func (s *shopServiceImpl) FetchShopByID(params *dto.ShopParams) (*domain.Shop, error) {
	if _, err := uuid.Parse(params.ID); err != nil {
		return nil, domain.ErrBadRequest
	}

	shop, err := s.shopRepo.FetchShopByID(params)

	return shop, err
}

func (s *shopServiceImpl) CreateShop(req *dto.ShopRequest) error {
	// image should be uploaded here!
	err := s.shopRepo.InsertShop(&domain.Shop{
		Name:        req.Name,
		Description: req.Description,
	})

	return err
}

func (s *shopServiceImpl) UpdateShop(req *dto.ShopRequest) error {
	err := s.shopRepo.UpdateShop(&domain.Shop{
		Name:        req.Name,
		Description: req.Description,
	})

	return err
}

func (s *shopServiceImpl) DeleteShop(params *dto.ShopParams) error {
	err := s.shopRepo.DeleteShop(params)

	return err
}
