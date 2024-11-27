package controller

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderSvc service.IOrderService
}

func MountOrderRoutes(r *gin.RouterGroup, orderSvc service.IOrderService) {
	orderCtr := &orderController{orderSvc}
	orderR := r.Group("/orders")

	orderR.GET("", orderCtr.FetchAll)
	orderR.GET("/:id", orderCtr.FetchByID)
	orderR.POST("", orderCtr.CreateOrder)
	orderR.PUT("/:id", orderCtr.UpdateOrder)
	orderR.DELETE("/:id", orderCtr.DeleteOrder)
}

// @Tags			Orders
// @Summary		Fetch All Orders
// @Description	Fetch All Orders From Database
// @Produce		json
// @Success		200	{object}	ginlib.Response{data=[]domain.Order}	"OK"
// @Failure		500	{object}	ginlib.Response							"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/orders [get]
func (c *orderController) FetchAll(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch all orders"
		orders  []domain.Order
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, orders, err)
	}()

	orders, err = c.orderSvc.FetchAllOrders()
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully fetch all orders"
}

// @Tags			Orders
// @Summary		Fetch Order By ID
// @Description	Fetch Order By ID From DB
// @Produce		json
// @Param			id	path		string								true	"Order ID"
// @Success		200	{object}	ginlib.Response{data=domain.Order}	"OK"
// @Failure		404	{object}	ginlib.Response						"Item not found"
// @Failure		500	{object}	ginlib.Response						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/orders/{id} [get]
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

// @Tags			Orders
// @Summary		Register Order
// @Description	Register Order to System
// @Produce		json
// @Param			OrderPayload	body		dto.OrderRequest	true	"Order Register Payload"
// @Success		200				{object}	ginlib.Response		"OK"
// @Failure		500				{object}	ginlib.Response		"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/orders [post]
func (c *orderController) CreateOrder(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch order"
		order   dto.OrderRequest
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBind(&order); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.orderSvc.CreateOrder(&dto.OrderParams{}, &order)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "succcessfully register order"

}

// @Tags			Orders
// @Summary		Update Order
// @Description	Update Existing Order
// @Produce		json
// @Param			OrderPayload	body		dto.OrderRequest	true	"Order Update Payload"
// @Param			id				path		string				true	"Order ID"
// @Success		200				{object}	ginlib.Response		"OK"
// @Failure		404				{object}	ginlib.Response		"Item not found"
// @Failure		500				{object}	ginlib.Response		"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/orders [put]
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

// @Tags			Orders
// @Summary		Delete Order
// @Description	Delete Existing Order from System
// @Produce		json
// @Param			id	path		string			true	"Order ID"
// @Success		200	{object}	ginlib.Response	"OK"
// @Failure		404	{object}	ginlib.Response	"Item not found"
// @Failure		500	{object}	ginlib.Response	"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/v1/orders [delete]
func (c *orderController) DeleteOrder(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to delete order"
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	err = c.orderSvc.DeleteOrder(&dto.OrderParams{
		ID: idParam,
	})
	code, status = domain.GetStatus(err)

	if err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	message = "successfully delete order"
}
