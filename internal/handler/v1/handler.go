package v1

import "github.com/labstack/echo/v4"

type Handler struct {
	link LinkService
}

func NewHandler(link LinkService) *Handler {
	return &Handler{
		link: link,
	}
}

func (h *Handler) Init(router *echo.Echo) {
	api := router.Group("/api/v1")
	{
		h.initLink(api)
	}
}
