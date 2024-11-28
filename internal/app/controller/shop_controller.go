package controller

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/devanfer02/filkom-canteen/internal/middleware"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

type shopController struct {
	shopSvc service.IShopService
}

func MountShopRoutes(r *gin.RouterGroup, shopSvc service.IShopService, mdlwr *middleware.Middleware) {
	shopCtr := &shopController{shopSvc}

	shopR := r.Group("/shops").Use(mdlwr.Authenticate())
	shopR.GET("", mdlwr.AuthorizeAdmin("Admin"), shopCtr.FetchAllShops)
	shopR.GET("/:id", mdlwr.AuthorizeAdmin("Admin", "Owner"), shopCtr.FetchShopByID)
	shopR.POST("", mdlwr.AuthorizeAdmin("Admin"), shopCtr.CreateShop)
	shopR.POST("/:id/owners/:ownerId", mdlwr.AuthorizeAdmin("Admin"),shopCtr.AssignOwner)
	shopR.DELETE("/:id/owners/:ownerId", mdlwr.AuthorizeAdmin("Admin"),shopCtr.RemoveOwner)
	shopR.PUT("/:id", mdlwr.AuthorizeAdmin("Admin", "Owner"), shopCtr.UpdateShop)
	shopR.DELETE("/:id", mdlwr.AuthorizeAdmin("Admin"), shopCtr.DeleteShop)
}

//	@Tags			Shops (Admin only)
//	@Summary		Fetch All Shops
//	@Description	Fetch All Shops From Database
//	@Produce		json
//	@Success		200	{object}	ginlib.Response{data=[]domain.Shop}	"OK"
//	@Failure		500	{object}	ginlib.Response						"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/shops [get]
func (c *shopController) FetchAllShops(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch all shops"
		shops   []domain.Shop
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, shops, err)
	}()

	shops, err = c.shopSvc.FetchAllShops()
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully fetch all shops"

}

//	@Tags			Shops (Admin and Owner)
//	@Summary		Fetch Shop By ID
//	@Description	Fetch Shop By ID From DB
//	@Produce		json
//	@Param			id	path		string								true	"Shop ID"
//	@Success		200	{object}	ginlib.Response{data=domain.Shop}	"OK"
//	@Failure		404	{object}	ginlib.Response{data=domain.Shop}	"Item not found"
//	@Failure		500	{object}	ginlib.Response						"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/shops/{id} [get]
func (c *shopController) FetchShopByID(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch all shops"
		shop    *domain.Shop
		err     error

		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, shop, err)
	}()

	shop, err = c.shopSvc.FetchShopByID(&dto.ShopParams{ID: idParam})
	code, status = domain.GetStatus(err)

	if err != nil {

		return
	}

	message = "successfully fetch shop by id"
}

//	@Tags			Shops (Admin only)
//	@Summary		Register Shop
//	@Description	Register Shop to System
//	@Produce		json
//	@Param			ShopPayload	body		dto.ShopRequest	true	"Shop Register Payload"
//	@Success		200			{object}	ginlib.Response	"OK"
//	@Failure		500			{object}	ginlib.Response	"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/shops [post]
func (c *shopController) CreateShop(ctx *gin.Context) {
	var (
		code    = 400
		status  = "fail"
		message = "failed to create new shop"
		shopReq dto.ShopRequest
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err := ctx.ShouldBind(&shopReq); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.shopSvc.CreateShop(&shopReq)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully create new shop"
}

//	@Tags			Shops (Admin only)
//	@Summary		Add Owner to Shop
//	@Description	Add Owner to Shop
//	@Produce		json
//	@Param			id		path		string			true	"Shop ID"
//	@Param			ownerId	path		string			true	"Owner ID"
//	@Success		200		{object}	ginlib.Response	"OK"
//	@Failure		404		{object}	ginlib.Response	"Shop or Owner not found"
//	@Failure		500		{object}	ginlib.Response	"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/shops/{id}/owners/{ownerId} [post]
func (c *shopController) AssignOwner(ctx *gin.Context) {
	var (
		code         = 400
		status       = "fail"
		message      = "failed to assign owner to shop"
		err          error
		idParam      = ctx.Param("id")
		ownerIdParam = ctx.Param("ownerId")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	err = c.shopSvc.AddOwner(&dto.ShopParams{
		ID:      idParam,
		OwnerID: ownerIdParam,
	})
	code, status = domain.GetStatus(err)

	if err != nil {
		return

	}

	message = "successfully assigned owner to shop"
}

//	@Tags			Shops (Admin only)
//	@Summary		Remove Owner from Shop
//	@Description	Remove Owner from Shop
//	@Produce		json
//	@Param			id		path		string			true	"Shop ID"
//	@Param			ownerId	path		string			true	"Owner ID"
//	@Success		200		{object}	ginlib.Response	"OK"
//	@Failure		404		{object}	ginlib.Response	"Item not found"
//	@Failure		500		{object}	ginlib.Response	"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/shops/{id}/owners/{ownerId} [delete]
func (c *shopController) RemoveOwner(ctx *gin.Context) {
	var (
		code         = 400
		status       = "fail"
		message      = "failed to remove owner from shop"
		err          error
		idParam      = ctx.Param("id")
		ownerIdParam = ctx.Param("ownerId")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	err = c.shopSvc.RemoveOwner(&dto.ShopParams{
		ID:      idParam,
		OwnerID: ownerIdParam,
	})
	code, status = domain.GetStatus(err)

	if err != nil {
		return

	}

	message = "successfully remove owner from shop"
}

//	@Tags			Shops (Admin and Owner)
//	@Summary		Update Shop
//	@Description	Update Existing Shop
//	@Produce		json
//	@Param			ShopPayload	body		dto.ShopRequest						true	"Shop Register Payload"
//	@Param			id			path		string								true	"Shop ID"
//	@Success		200			{object}	ginlib.Response						"OK"
//	@Failure		404			{object}	ginlib.Response{data=domain.Shop}	"Item not found"
//	@Failure		409			{object}	ginlib.Response						"Username already exists"
//	@Failure		500			{object}	ginlib.Response						"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/shops/{id} [put]
func (c *shopController) UpdateShop(ctx *gin.Context) {
	var (
		code    = 400
		status  = "fail"
		message = "failed to update shop"
		shopReq dto.ShopRequest
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err := ctx.ShouldBind(&shopReq); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.shopSvc.UpdateShop(&dto.ShopParams{ID: idParam}, &shopReq)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully update shop"
}

//	@Tags			Shops (Admin only)
//	@Summary		Delete Shop
//	@Description	Delete Existing Shop
//	@Produce		json
//	@Param			id	path		string			true	"Shop ID"
//	@Success		200	{object}	ginlib.Response	"OK"
//	@Failure		404	{object}	ginlib.Response	"Item not found"
//	@Failure		500	{object}	ginlib.Response	"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/shops/{id} [delete]
func (c *shopController) DeleteShop(ctx *gin.Context) {
	var (
		code    = 400
		status  = "fail"
		message = "failed to delete shop"
		shopReq dto.ShopRequest
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err := ctx.ShouldBind(&shopReq); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.shopSvc.DeleteShop(&dto.ShopParams{ID: idParam})
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully delete shop"
}
