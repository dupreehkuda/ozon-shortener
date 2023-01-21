package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router returns service's echo router
func (s server) Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/:id", s.httpHandlers.GetFullLink)
	e.POST("/api/shorten", s.httpHandlers.ShortenLink)

	return e
}
