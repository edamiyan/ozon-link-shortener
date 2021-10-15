package v1

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) initLink(api *echo.Group) {
	links := api.Group("/tokens")
	{
		links.GET("/:token", h.getBase)
		links.POST("", h.createShort)
	}
}

func (h *Handler) createShort(ctx echo.Context) error {
	panic("implement me")
}

func (h *Handler) getBase(ctx echo.Context) error {
	panic("implement me")
}
