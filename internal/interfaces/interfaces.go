package interfaces

import "github.com/labstack/echo/v4"

//go:generate mockgen -source interfaces.go -destination ../mock/mock.go -package mock;

type Handlers interface {
	GetFullLink(c echo.Context) error
	ShortenLink(c echo.Context) error
}

type Service interface {
	GetFullLink(id string) (string, error)
	ShortenLink(link string) (string, error)
}

type Storage interface {
	AddNewLink(id, link string) error
	GetFullLink(id string) (string, error)
}
