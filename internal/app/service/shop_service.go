package service

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	enc "github.com/devanfer02/filkom-canteen/internal/pkg/encoder"
	"github.com/google/uuid"
)

type IShopService interface {
	FetchAllShops() ([]domain.Shop, error)
	FetchShopByID(params *dto.ShopParams) (*domain.Shop, error)
	CreateShop(req *dto.ShopRequest) error
	AddOwner(req *dto.ShopParams) error
	RemoveOwner(req *dto.ShopParams) error
	UpdateShop(params *dto.ShopParams, req *dto.ShopRequest) error
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

	if err != nil {
		return nil, err 
	}

	for idx, shop := range shops {
		shops[idx].ID = enc.Encode(shop.ID)
	}

	return shops, err
}

func (s *shopServiceImpl) FetchShopByID(params *dto.ShopParams) (*domain.Shop, error) {
	decoded, err := enc.Decode(params.ID)

	if err != nil {
		return nil, domain.ErrBadRequest
	}

	params.ID = decoded

	if _, err := uuid.Parse(params.ID); err != nil {
		return nil, domain.ErrBadRequest
	}

	shop, err := s.shopRepo.FetchShopByID(params)

	if err != nil {
		return nil, err 
	}

	shop.ID = enc.Encode(shop.ID) 

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

func (s *shopServiceImpl) AddOwner(req *dto.ShopParams) error {
	err := s.shopRepo.InsertShopOwner(req)

	return err
}

func (s *shopServiceImpl) RemoveOwner(req *dto.ShopParams) error {
	err := s.shopRepo.DeleteShopOwner(req)

	return err
}

func (s *shopServiceImpl) UpdateShop(params *dto.ShopParams, req *dto.ShopRequest) error {
	decoded, err := enc.Decode(params.ID)

	if err != nil {
		return domain.ErrBadRequest
	}

	params.ID = decoded

	if _, err := uuid.Parse(params.ID); err != nil {
		return domain.ErrBadRequest
	}

	err = s.shopRepo.UpdateShop(params, &domain.Shop{
		Name:        req.Name,
		Description: req.Description,
	})

	return err
}

func (s *shopServiceImpl) DeleteShop(params *dto.ShopParams) error {
	decoded, err := enc.Decode(params.ID)

	if err != nil {
		return domain.ErrBadRequest
	}

	params.ID = decoded

	if _, err := uuid.Parse(params.ID); err != nil {
		return domain.ErrBadRequest
	}

	err = s.shopRepo.DeleteShop(params)

	return err
}
