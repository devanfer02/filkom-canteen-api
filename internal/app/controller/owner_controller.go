package controller

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/devanfer02/filkom-canteen/internal/middleware"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

type ownerController struct {
	ownerSvc service.IOwnerService
}

func MountOwnerRoutes(r *gin.RouterGroup, ownerSvc service.IOwnerService, mdlwr *middleware.Middleware) {
	ownerCtr := &ownerController{ownerSvc}
	ownerR := r.Group("/owners")

	ownerR.GET("", mdlwr.Authenticate(), mdlwr.AuthorizeAdmin("Admin"), ownerCtr.FetchAll)
	ownerR.GET("/:id", mdlwr.Authenticate(), mdlwr.AuthorizeAdmin("Admin", "Owner"), ownerCtr.FetchByID)
	ownerR.POST("", mdlwr.Authenticate(), mdlwr.AuthorizeAdmin("Admin"), ownerCtr.RegisterOwner)
	ownerR.PUT("/:id", mdlwr.Authenticate(), mdlwr.AuthorizeAdmin("Admin", "Owner"), ownerCtr.UpdateOwner)
	ownerR.DELETE("/:id", mdlwr.Authenticate(), mdlwr.AuthorizeAdmin("Admin"), ownerCtr.DeleteOwner)
}

// @Tags			Owners
// @Summary		Fetch All Owners
// @Description	Fetch All Owners From Database
// @Produce		json
// @Success		200	{object}	ginlib.Response{data=[]domain.Owner}	"OK"
// @Failure		500	{object}	ginlib.Response							"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/owners [get]
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

// @Tags			Owners
// @Summary		Fetch Owner By ID
// @Description	Fetch Owner By ID From DB
// @Produce		json
// @Param			id	path		string								true	"Owner ID"
// @Success		200	{object}	ginlib.Response{data=domain.Owner}	"OK"
// @Failure		404	{object}	ginlib.Response{data=domain.Owner}	"Item not found"
// @Failure		500	{object}	ginlib.Response						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/owners/{id} [get]
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

// @Tags			Owners
// @Summary		Register Owner
// @Description	Register Owner to System
// @Produce		json
// @Param			OwnerPayload	body		dto.OwnerRequest	true	"Owner Register Payload"
// @Success		200				{object}	ginlib.Response		"OK"
// @Failure		409				{object}	ginlib.Response		"Username already exists"
// @Failure		500				{object}	ginlib.Response		"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/owners [post]
func (c *ownerController) RegisterOwner(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to register owner"
		owner   dto.OwnerRequest
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBindJSON(&owner); err != nil {
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

// @Tags			Owners
// @Summary		Update Owner
// @Description	Update Existing Owner
// @Produce		json
// @Param			OwnerPayload	body		dto.OwnerRequest					true	"Owner Register Payload"
// @Param			id				path		string								true	"Owner ID"
// @Success		200				{object}	ginlib.Response						"OK"
// @Failure		404				{object}	ginlib.Response{data=domain.Owner}	"Item not found"
// @Failure		409				{object}	ginlib.Response						"Username already exists"
// @Failure		500				{object}	ginlib.Response						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/owners/{id} [put]
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

	if err = ctx.ShouldBindJSON(&owner); err != nil {
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

// @Tags			Owners
// @Summary		Delete Owner
// @Description	Delete Existing Owner
// @Produce		json
// @Param			id	path		string			true	"Owner ID"
// @Success		200	{object}	ginlib.Response	"OK"
// @Failure		404	{object}	ginlib.Response	"Item not found"
// @Failure		500	{object}	ginlib.Response	"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/owners/{id} [delete]
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
