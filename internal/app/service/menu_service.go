package service

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/google/uuid"
)

type IMenuService interface {
	FetchAllMenus() ([]domain.Menu, error)
	FetchMenuByID(params *dto.MenuParams) (*domain.Menu, error)
	CreateMenu(req *dto.MenuRequest) error
	UpdateMenu(params *dto.MenuParams,req *dto.MenuRequest) error
	DeleteMenu(params *dto.MenuParams) error	
}

type menuServiceImpl struct {
	menuRepo repository.IMenuRepository
}

func NewMenuService(menuRepo repository.IMenuRepository) IMenuService {
	return &menuServiceImpl{menuRepo}
}

func(s *menuServiceImpl) FetchAllMenus() ([]domain.Menu, error) {
	menus, err := s.menuRepo.FetchAll(&dto.MenuParams{})

	return menus, err 
}

func(s *menuServiceImpl) FetchMenuByID(params *dto.MenuParams) (*domain.Menu, error) {
	if _, err := uuid.Parse(params.ID); err != nil {
		return nil, domain.ErrBadRequest
	}

	menu, err := s.menuRepo.FetchByID(params)

	return menu, err 
}

func(s *menuServiceImpl) CreateMenu(req *dto.MenuRequest) error {
	err := s.menuRepo.InsertMenu(&domain.Menu{
		Name: req.Name,
		ShopID: req.ShopID,
		Price: req.Price,
		Status: req.Status,
	})

	return err 
}

func(s *menuServiceImpl) UpdateMenu(params *dto.MenuParams,req *dto.MenuRequest) error {

	err := s.menuRepo.UpdateMenu(params, &domain.Menu{
		Name: req.Name,
		ShopID: req.ShopID,
		Price: req.Price,
		Status: req.Status,
	})

	return err 
}

func(s *menuServiceImpl) DeleteMenu(params *dto.MenuParams) error {
	err := s.menuRepo.DeleteMenu(params)

	return err 
}