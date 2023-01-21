package httpHandlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
	validators "github.com/dupreehkuda/ozon-shortener/internal/handlers"
	"github.com/dupreehkuda/ozon-shortener/internal/models"
)

// ShortenLink handles the operation of shortening new link
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
		if err == er.ErrExistingToken {
			h.logger.Error("Collision alert!", zap.Error(err))
			return c.JSON(http.StatusConflict, nil)
		}

		h.logger.Error("Error occurred in service call", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	link := fmt.Sprintf("http://%s/%s", c.Request().Host, token)

	return c.JSON(http.StatusOK, models.ShortenResponse{Link: link})
}
