package grpcHandlers

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
	validators "github.com/dupreehkuda/ozon-shortener/internal/handlers"
	rpcapi "github.com/dupreehkuda/ozon-shortener/pkg/api"
)

// ShortenLink handles the operation of shortening new link
func (h handlers) ShortenLink(ctx context.Context, req *rpcapi.ShortenRequest) (*rpcapi.ShortenResponse, error) {
	if err := validators.ValidateURL(req.Url); err != nil {
		h.logger.Debug("Validation", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "Validation error")
	}

	token, err := h.service.ShortenLink(req.Url)
	if err != nil {
		if err == er.ErrExistingToken {
			h.logger.Error("Collision alert!", zap.Error(err))
			return nil, status.Error(codes.AlreadyExists, "Collision alert!")
		}

		h.logger.Error("Error occurred in service call", zap.Error(err))
		return nil, err
	}

	return &rpcapi.ShortenResponse{Token: token}, nil
}
