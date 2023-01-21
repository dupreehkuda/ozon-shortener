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

// GetFullLink handles the operation of redirect on shortened
func (h handlers) GetFullLink(ctx context.Context, req *rpcapi.GetFullRequest) (*rpcapi.GetFullResponse, error) {
	if err := validators.ValidateToken(req.Token); err != nil {
		h.logger.Debug("Validation", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "Validation error")
	}

	link, err := h.service.GetFullLink(req.Token)
	if err != nil {
		if err == er.ErrNoSuchURL {
			return nil, status.Error(codes.NotFound, "URL not found")
		}

		h.logger.Error("Error occurred in service call", zap.Error(err))
		return nil, err
	}

	return &rpcapi.GetFullResponse{Url: link}, nil
}
