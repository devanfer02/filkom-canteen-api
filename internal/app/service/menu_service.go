package service

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	enc "github.com/devanfer02/filkom-canteen/internal/pkg/encoder"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
	"github.com/google/uuid"
)

type IMenuService interface {
	FetchAllMenus(params *dto.MenuParams) ([]domain.Menu, error)
	FetchMenuByID(params *dto.MenuParams) (*domain.Menu, error)
	CreateMenu(params *dto.MenuParams, req *dto.MenuRequest) error
	UpdateMenu(params *dto.MenuParams, req *dto.MenuRequest) error
	DeleteMenu(params *dto.MenuParams) error
}

type menuServiceImpl struct {
	menuRepo repository.IMenuRepository
}

func NewMenuService(menuRepo repository.IMenuRepository) IMenuService {
	return &menuServiceImpl{menuRepo}
}

func (s *menuServiceImpl) FetchAllMenus(params *dto.MenuParams) ([]domain.Menu, error) {
	menus, err := s.menuRepo.FetchAll(params)

	if err != nil {
		return nil, err 
	}

	for idx, menu := range menus {
		menus[idx].ID = enc.Encode(menu.ID)
		menus[idx].ShopID = enc.Encode(menu.ShopID)
	}

	return menus, err
}

func (s *menuServiceImpl) FetchMenuByID(params *dto.MenuParams) (*domain.Menu, error) {
	decoded, err := enc.Decode(params.ID)

	if err != nil {
		return nil, domain.ErrBadRequest
	}

	params.ID = decoded

	if _, err := uuid.Parse(params.ID); err != nil {
		return nil, domain.ErrBadRequest
	}

	menu, err := s.menuRepo.FetchByID(params)

	if err != nil {
		return nil, err 
	}

	menu.ID = enc.Encode(menu.ID) 
	menu.ShopID = enc.Encode(menu.ShopID)

	log.Info(log.LogInfo{
		"LTE": menu.ID,
	}, "OK")

	return menu, err
}

func (s *menuServiceImpl) CreateMenu(params *dto.MenuParams, req *dto.MenuRequest) error {
	decodedShopID, err := enc.Decode(req.ShopID)

	if err != nil {
		return domain.ErrBadRequest
	}

	err = s.menuRepo.InsertMenu(&domain.Menu{
		Name:   req.Name,
		ShopID: decodedShopID,
		Price:  req.Price,
		Status: req.Status,
	})

	return err
}

func (s *menuServiceImpl) UpdateMenu(params *dto.MenuParams, req *dto.MenuRequest) error {

	err := s.menuRepo.UpdateMenu(params, &domain.Menu{
		Name:   req.Name,
		ShopID: req.ShopID,
		Price:  req.Price,
		Status: req.Status,
	})

	return err
}

func (s *menuServiceImpl) DeleteMenu(params *dto.MenuParams) error {
	err := s.menuRepo.DeleteMenu(params)

	return err
}
