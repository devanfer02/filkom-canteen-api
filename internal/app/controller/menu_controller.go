package controller

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/app/service"
	"github.com/devanfer02/filkom-canteen/internal/dto"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

type menuController struct {
	menuSvc service.IMenuService
}

func MountMenuRoutes(r *gin.RouterGroup, menuSvc service.IMenuService) {
	menuCtr := &menuController{menuSvc}
	menuR := r.Group("/menus")

	menuR.GET("", menuCtr.FetchAll)
	menuR.GET("/:id", menuCtr.FetchByID)
	menuR.POST("", menuCtr.CreateMenu)
	menuR.PUT("/:id", menuCtr.UpdateMenu)
	menuR.DELETE("/:id", menuCtr.DeleteMenu)
}

//	@Tags			Menus
//	@Summary		Fetch All Menus
//	@Description	Fetch All Menus From Database
//	@Produce		json
//	@Success		200	{object}	ginlib.Response{data=[]domain.Menu}	"OK"
//	@Failure		500	{object}	ginlib.Response							"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/menus [get]
func (c *menuController) FetchAll(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch all menus"
		menus   []domain.Menu
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, menus, err)
	}()

	menus, err = c.menuSvc.FetchAllMenus()
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully fetch all menus"
}

//	@Tags			Menus
//	@Summary		Fetch Menu By ID
//	@Description	Fetch Menu By ID From DB
//	@Produce		json
//	@Param			id	path		string								true	"Menu ID"
//	@Success		200	{object}	ginlib.Response{data=domain.Menu}	"OK"
//	@Failure		404	{object}	ginlib.Response{data=domain.Menu}	"Item not found"
//	@Failure		500	{object}	ginlib.Response						"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/menus/{id} [get]
func (c *menuController) FetchByID(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch menu"
		menu    *domain.Menu
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, menu, err)
	}()

	menu, err = c.menuSvc.FetchMenuByID(&dto.MenuParams{ID: idParam})
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully fetch menu"
}

//	@Tags			Menus
//	@Summary		Register Menu
//	@Description	Register Menu to System
//	@Produce		json
//	@Param			MenuPayload	body		dto.MenuRequest	true	"Menu Register Payload"
//	@Success		200				{object}	ginlib.Response		"OK"
//	@Failure		500				{object}	ginlib.Response		"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/menus [post]
func (c *menuController) CreateMenu(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch menu"
		menu    dto.MenuRequest
		err     error
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBind(&menu); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.menuSvc.CreateMenu(&dto.MenuParams{}, &menu)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "succcessfully register menu"

}

//	@Tags			Menus
//	@Summary		Update Menu
//	@Description	Update Existing Menu
//	@Produce		json
//	@Param			MenuPayload	body		dto.MenuRequest	true	"Menu Update Payload"
//	@Param			id				path		string								true	"Menu ID"
//	@Success		200				{object}	ginlib.Response		"OK"
// 	@Failure		404				{object}	ginlib.Response		"Item not found"
//	@Failure		500				{object}	ginlib.Response		"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/menus [put]
func (c *menuController) UpdateMenu(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to fetch menu"
		menu    dto.MenuRequest
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	if err = ctx.ShouldBind(&menu); err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	err = c.menuSvc.UpdateMenu(&dto.MenuParams{
		ID: idParam,
	}, &menu)
	code, status = domain.GetStatus(err)

	if err != nil {
		return
	}

	message = "successfully update menu"
}

//	@Tags			Menus
//	@Summary		Delete Menu
//	@Description	Delete Existing Menu from System
//	@Produce		json
//	@Param			id				path		string								true	"Menu ID"
//	@Success		200				{object}	ginlib.Response		"OK"
// 	@Failure		404				{object}	ginlib.Response		"Item not found"
//	@Failure		500				{object}	ginlib.Response		"Internal Server Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/menus [delete]
func (c *menuController) DeleteMenu(ctx *gin.Context) {
	var (
		code    = 500
		status  = "fail"
		message = "failed to delete menu"
		err     error
		idParam = ctx.Param("id")
	)

	defer func() {
		ginlib.SendResponse(ctx, code, status, message, nil, err)
	}()

	err = c.menuSvc.DeleteMenu(&dto.MenuParams{
		ID: idParam,
	})
	code, status = domain.GetStatus(err)

	if err != nil {
		code, status = domain.GetStatus(err)
		return
	}

	message = "successfully delete menu"
}
