package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	validators "github.com/dupreehkuda/ozon-shortener/internal/handlers"
	"github.com/dupreehkuda/ozon-shortener/internal/models"
)

func (h handlers) ShortenLink(c echo.Context) error {
	req := &models.ShortenRequest{}

	if err := c.Bind(req); err != nil {
		h.logger.Error("Error occurred binding value", zap.Error(err))
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := validators.ValidateURL(req.URL); err != nil {
		h.logger.Debug("Validation", zap.Error(err))
		return c.JSON(http.StatusBadRequest, nil)
	}

	token, err := h.service.ShortenLink(req.URL)
	if err != nil {
		h.logger.Error("Error occurred in service call", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	link := fmt.Sprintf("http://%s/%s", c.Request().Host, token)

	return c.JSON(http.StatusOK, models.ShortenResponse{Link: link})
}
