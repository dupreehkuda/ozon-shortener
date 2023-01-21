package interfaces

import (
	"github.com/labstack/echo/v4"

	rpcapi "github.com/dupreehkuda/ozon-shortener/pkg/api"
)

//go:generate mockgen -source interfaces.go -destination ../mock/mock.go -package mock;

// RPCHandlers is interface for gRPC handlers
type RPCHandlers interface {
	rpcapi.ShortenerServer
}

// Handlers is interface for http handlers
type Handlers interface {
	GetFullLink(c echo.Context) error
	ShortenLink(c echo.Context) error
}

// Service is interface for business-logic
type Service interface {
	GetFullLink(id string) (string, error)
	ShortenLink(link string) (string, error)
}

// Storage is interface for storage
type Storage interface {
	AddNewLink(id, link string) (string, error)
	GetFullLink(id string) (string, error)
}
