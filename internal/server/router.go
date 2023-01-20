package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s server) router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/:id", nil)
	e.POST("/", nil)

	return e
}
