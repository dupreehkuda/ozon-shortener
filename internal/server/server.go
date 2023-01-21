package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/dupreehkuda/ozon-shortener/internal/config"
	i "github.com/dupreehkuda/ozon-shortener/internal/interfaces"
	rpcapi "github.com/dupreehkuda/ozon-shortener/pkg/api"
)

// server provides single configuration out of all components
type server struct {
	httpHandlers i.Handlers
	grpcHandlers i.RPCHandlers
	config       *config.Config
	logger       *zap.Logger
}

// NewByConfig returns server instance with default config
func NewByConfig(httpHandlers i.Handlers, grpcHandlers i.RPCHandlers, logger *zap.Logger, config *config.Config) *server {
	return &server{
		httpHandlers: httpHandlers,
		grpcHandlers: grpcHandlers,
		logger:       logger,
		config:       config,
	}
}

// Run runs the service and provides graceful shutdown
func (s server) Run() {
	httpServ := &http.Server{Addr: s.config.HTTPAddress, Handler: s.Router()}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	go s.RunGRPC(serverCtx, sig)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				s.logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := httpServ.Shutdown(shutdownCtx)
		if err != nil {
			s.logger.Fatal("Error shutting down", zap.Error(err))
		}

		s.logger.Info("Server shut down", zap.String("http", s.config.HTTPAddress), zap.String("grpc", s.config.GRPCAddress))
		serverStopCtx()
	}()

	s.logger.Info("Server started", zap.String("http", s.config.HTTPAddress), zap.String("grpc", s.config.GRPCAddress))

	// http server start
	if err := httpServ.ListenAndServe(); err != nil {
		s.logger.Fatal("Cant start server", zap.Error(err))
	}

	<-serverCtx.Done()
}

// RunGRPC runs gRPC server with a separate shutdown goroutine
func (s server) RunGRPC(ctx context.Context, sig chan os.Signal) {
	grpcServ := grpc.NewServer()
	rpcapi.RegisterShortenerServer(grpcServ, s.grpcHandlers)

	grpcListener, err := net.Listen("tcp", s.config.GRPCAddress)
	if err != nil {
		s.logger.Fatal("server down")
	}

	// grpc server start
	if err = grpcServ.Serve(grpcListener); err != nil {
		s.logger.Fatal("Cant start server", zap.Error(err))
	}

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				s.logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		grpcServ.GracefulStop()
	}()
}
