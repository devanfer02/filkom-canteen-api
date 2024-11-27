package service

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/devanfer02/filkom-canteen/internal/pkg/bcrypt"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
	"github.com/google/uuid"
)

type IOwnerService interface {
	FetchAllOwners() ([]domain.Owner, error)
	FetchOwnerByID(params *dto.OwnerParams) (*domain.Owner, error)
	CreateOwner(req *dto.OwnerRequest) error
	UpdateOwner(params *dto.OwnerParams, req *dto.OwnerRequest) error
	DeleteOwner(params *dto.OwnerParams) error
}

type ownerServiceImpl struct {
	ownerRepo repository.IOwnerRepository
}

func NewOwnerService(ownerRepo repository.IOwnerRepository) IOwnerService {
	return &ownerServiceImpl{ownerRepo}
}

func (s *ownerServiceImpl) FetchAllOwners() ([]domain.Owner, error) {
	owners, err := s.ownerRepo.FetchAll(&dto.OwnerParams{})

	return owners, err
}

func (s *ownerServiceImpl) FetchOwnerByID(params *dto.OwnerParams) (*domain.Owner, error) {
	if _, err := uuid.Parse(params.ID); err != nil {
		return nil, domain.ErrBadRequest
	}

	owner, err := s.ownerRepo.FetchByID(params)

	return owner, err
}

func (s *ownerServiceImpl) CreateOwner(req *dto.OwnerRequest) error {
	var (
		err error
	)

	req.Password, err = bcrypt.HashPassword(req.Password)

	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[BLOG SERVICE][CreateOwner] failed to create owner")
		return err
	}

	err = s.ownerRepo.InsertOwner(&domain.Owner{
		Fullname: req.Fullname,
		Username: req.Username,
		Password: req.Password,
		WANumber: req.WANumber,
	})

	return err
}

func (s *ownerServiceImpl) UpdateOwner(params *dto.OwnerParams, req *dto.OwnerRequest) error {
	var (
		err error
	)

	if req.Password != "" {
		req.Password, err = bcrypt.HashPassword(req.Password)

		if err != nil {
			log.Error(log.LogInfo{
				"error": err.Error(),
			}, "[BLOG SERVICE][CreateOwner] failed to create owner")
			return err
		}
	}

	err = s.ownerRepo.UpdateOwner(params, &domain.Owner{
		Fullname: req.Fullname,
		Username: req.Username,
		Password: req.Password,
		WANumber: req.WANumber,
	})

	return nil
}

func (s *ownerServiceImpl) DeleteOwner(params *dto.OwnerParams) error {
	err := s.ownerRepo.DeleteOwner(params)

	return err
}
