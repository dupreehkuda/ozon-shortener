package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handlers) GetShortenedLink(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": "test",
	})
}
