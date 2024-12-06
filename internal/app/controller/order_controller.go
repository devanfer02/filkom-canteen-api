package controller

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	"github.com/devanfer02/filkom-canteen/internal/middleware"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderSvc service.IOrderService
}

func MountOrderRoutes(r *gin.RouterGroup, orderSvc service.IOrderService, mdlwr *middleware.Middleware) {
	orderCtr := &orderController{orderSvc}
	orderR := r.Group("/orders")

	orderR.GET("", mdlwr.Authenticate(),orderCtr.FetchAll)
	orderR.GET("/:id", mdlwr.Authenticate(), orderCtr.FetchByID)
	orderR.POST("", mdlwr.Authenticate(),mdlwr.RateLimiter(50),  orderCtr.CreateOrder)
	orderR.PUT("/:id", mdlwr.Authenticate(), mdlwr.AuthorizeAdmin("Admin", "Owner"),orderCtr.UpdateOrder)
	orderR.DELETE("/:id", mdlwr.Authenticate(), orderCtr.DeleteOrder)
}

//	@Tags			Orders
//	@Summary		Fetch All Orders
//	@Description	Fetch All Orders From Database
//	@Produce		json
//	@Param			shop_id	query		string									false	"Shop ID"
//	@Success		200		{object}	ginlib.Response{data=[]domain.Order}	"OK"
//	@Failure		500		{object}	ginlib.Response							"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/orders [get]
func (c *orderController) FetchAll(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch all orders"
		orders  []domain.Order
		err     error
		shopId  = ctx.Query("shop_id")
		userID  = ctx.GetString("id")
		user    = ctx.GetString("user")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, orders, err)
	}()

	orders, err = c.orderSvc.FetchAllOrders(&dto.OrderParams{
		ShopID: shopId,
		UserID: func()string {
			if user == env.AppEnv.JWTUserRole {
				return userID 
			}

			return ""
		}(),
	})

	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully fetch all orders"
}

//	@Tags			Orders
//	@Summary		Fetch Order By ID
//	@Description	Fetch Order By ID From DB
//	@Produce		json
//	@Param			id	path		string								true	"Order ID"
//	@Success		200	{object}	ginlib.Response{data=domain.Order}	"OK"
//	@Failure		404	{object}	ginlib.Response						"Item not found"
//	@Failure		500	{object}	ginlib.Response						"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/orders/{id} [get]
func (c *orderController) FetchByID(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch order"
		order   *domain.Order
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, order, err)
	}()

	order, err = c.orderSvc.FetchOrderByID(&dto.OrderParams{ID: idParam})
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully fetch order"
}

//	@Tags			Orders
//	@Summary		Register Order
//	@Description	Register Order to System
//	@Produce		json
//	@Param			OrderPayload	body		dto.OrderRequest	true	"Order Register Payload"
//	@Success		200				{object}	ginlib.Response		"OK"
//	@Failure		500				{object}	ginlib.Response		"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/orders [post]
func (c *orderController) CreateOrder(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch order"
		order   dto.OrderRequest
		err     error
		userId  = ctx.GetString("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBind(&order); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.orderSvc.CreateOrder(&dto.OrderParams{
		UserID: userId,
	}, &order)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "succcessfully register order"

}

//	@Tags			Orders (Admin and Owner)
//	@Summary		Update Order
//	@Description	Update Existing Order
//	@Produce		json
//	@Param			OrderPayload	body		dto.OrderRequest	true	"Order Update Payload"
//	@Param			id				path		string				true	"Order ID"
//	@Success		200				{object}	ginlib.Response		"OK"
//	@Failure		404				{object}	ginlib.Response		"Item not found"
//	@Failure		500				{object}	ginlib.Response		"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/orders [put]
func (c *orderController) UpdateOrder(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch order"
		order   dto.OrderRequest
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBind(&order); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.orderSvc.UpdateOrder(&dto.OrderParams{
		ID: idParam,
	}, &order)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully update order"
}

//	@Tags			Orders
//	@Summary		Delete Order
//	@Description	Delete Existing Order from System
//	@Produce		json
//	@Param			id	path		string			true	"Order ID"
//	@Success		200	{object}	ginlib.Response	"OK"
//	@Failure		404	{object}	ginlib.Response	"Item not found"
//	@Failure		500	{object}	ginlib.Response	"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Security		UserAuth
//	@Router			/api/v1/orders [delete]
func (c *orderController) DeleteOrder(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to delete order"
		err     error
		idParam = ctx.Param("id")
		userId  = ctx.GetString("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	err = c.orderSvc.DeleteOrder(&dto.OrderParams{
		ID: idParam,
		UserID: userId,
	})
	code, status = domain.GetStatus(err)

	if err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	message = "successfully delete order"
}
