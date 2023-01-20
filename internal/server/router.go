package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s server) Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/:id", s.handlers.GetFullLink)
	e.POST("/", s.handlers.ShortenLink)

	return e
}
