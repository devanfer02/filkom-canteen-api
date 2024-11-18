package controller

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

type shopController struct {
	shopSvc service.IShopService
}

func MountShopRoutes(r *gin.RouterGroup, shopSvc service.IShopService) {
	shopCtr := &shopController{shopSvc}

	shopR := r.Group("/shops")
	shopR.GET("", shopCtr.FetchAllShops)
	shopR.GET("/:id", shopCtr.FetchShopByID)
	shopR.POST("", shopCtr.CreateShop)
	shopR.PUT("/:id", shopCtr.UpdateShop)
	shopR.DELETE("/:id", shopCtr.DeleteShop)
}

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

func (c *shopController) CreateShop(ctx *gin.Context) {
	var (
		code = 400
		status = "fail"
		message = "failed to create new shop"
		shopReq dto.ShopRequest
		err error 
	)

	defer func(){
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err := ctx.ShouldBind(&shopReq); err != nil {
		return 
	}

	err = c.shopSvc.CreateShop(&shopReq)
	code, status = domain.GetStatus(err)

	if err != nil {
		return 
	}

	message = "successfully create new shop"
}

func (c *shopController) UpdateShop(ctx *gin.Context) {

}

func (c *shopController) DeleteShop(ctx *gin.Context) {

}
