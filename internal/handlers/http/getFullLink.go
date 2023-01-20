package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
	validators "github.com/dupreehkuda/ozon-shortener/internal/handlers"
)

func (h handlers) GetFullLink(c echo.Context) error {
	token := c.Param("id")

	if err := validators.ValidateToken(token); err != nil {
		h.logger.Debug("Validation", zap.Error(err))
		return c.JSON(http.StatusBadRequest, nil)
	}

	link, err := h.service.GetFullLink(token)
	if err != nil {
		if err == er.ErrNoSuchURL {
			return c.JSON(http.StatusNoContent, nil)
		}

		h.logger.Error("Error occurred in service call", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.Redirect(http.StatusMovedPermanently, link)
}
