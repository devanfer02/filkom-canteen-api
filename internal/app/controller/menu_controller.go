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
