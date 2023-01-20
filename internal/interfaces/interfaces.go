package interfaces

import "github.com/labstack/echo/v4"

type Handlers interface {
	GetShortenedLink(c echo.Context) error
	ShortenLink(c echo.Context) error
}

type Service interface {
	GetShortenedLink(id string) (string, error)
	ShortenLink(link string) (string, error)
}

type Storage interface {
}
