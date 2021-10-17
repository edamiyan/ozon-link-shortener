package v1

import (
	"github.com/edamiyan/ozon-link-shortener/internal/model"
	"github.com/edamiyan/ozon-link-shortener/internal/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) initLink(api *echo.Group) {
	links := api.Group("/tokens")
	{
		links.GET("/:token", h.getBase)
		links.POST("", h.createShort)
	}
}

func (h *Handler) createShort(ctx echo.Context) error {
	input := &model.Link{}
	if err := ctx.Bind(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewValidationError("can't bind input link data"))
	}

	if err := model.ValidateBaseURL(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewValidationError("validation error"))
	}

	token, err := h.link.CreateShortURL(ctx.Request().Context(), input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.NewError("something went wrong"))
	}

	return ctx.JSON(http.StatusOK, token)
}

func (h *Handler) getBase(ctx echo.Context) error {
	panic("implement me")
}
