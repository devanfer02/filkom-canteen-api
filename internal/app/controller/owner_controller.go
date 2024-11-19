package controller

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

type ownerController struct {
	ownerSvc service.IOwnerService
}

func MountOwnerRoutes(r *gin.RouterGroup, ownerSvc service.IOwnerService) {
	ownerCtr := &ownerController{ownerSvc}
	ownerR := r.Group("/owners")

	ownerR.GET("", ownerCtr.FetchAll)
	ownerR.GET("/:id", ownerCtr.FetchByID)
	ownerR.POST("", ownerCtr.RegisterOwner)
	ownerR.PUT("/:id", ownerCtr.UpdateOwner)
	ownerR.DELETE("/:id", ownerCtr.DeleteOwner)
}

func (c *ownerController) FetchAll(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch all owners"
		owners  []domain.Owner
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, owners, err)
	}()

	owners, err = c.ownerSvc.FetchAllOwners()
	code, status = domain.GetStatus(err)

	if err != nil {
		return 
	}

	message = "successfully fetch all owners"
}

func (c *ownerController) FetchByID(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch owner"
		owner   *domain.Owner
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, owner, err)
	}()

	owner, err = c.ownerSvc.FetchOwnerByID(&dto.OwnerParams{ID: idParam})
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully fetch owner"
}

func (c *ownerController) RegisterOwner(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch owner"
		owner   dto.OwnerRequest
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBind(&owner); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.ownerSvc.CreateOwner(&owner)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "succcessfully register owner"

}

func (c *ownerController) UpdateOwner(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch owner"
		owner   dto.OwnerRequest
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBind(&owner); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.ownerSvc.UpdateOwner(&dto.OwnerParams{
		ID: idParam,
	}, &owner)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully update owner"
}

func (c *ownerController) DeleteOwner(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to delete owner"
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	err = c.ownerSvc.DeleteOwner(&dto.OwnerParams{
		ID: idParam,
	})
	code, status = domain.GetStatus(err)

	if err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	message = "successfully delete owner"
}
