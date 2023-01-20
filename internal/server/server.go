package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/ozon-shortener/internal/config"
	i "github.com/dupreehkuda/ozon-shortener/internal/interfaces"
	"github.com/dupreehkuda/ozon-shortener/internal/logger"
)

// server provides single configuration out of all components
type server struct {
	handlers i.Handlers
	config   *config.Config
	logger   *zap.Logger
}

// NewByConfig returns server instance with default config
func NewByConfig() *server {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	//logic := processors.New(uc, tc, log)

	//handle := handlers.New(logic, log)

	return &server{
		//handlers:   handle,
		//middleware: mv,
		logger: log,
		config: cfg,
	}
}

// Run runs the service and provides graceful shutdown
func (s server) Run() {
	serv := &http.Server{Addr: s.config.Address, Handler: s.router()}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
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

		err := serv.Shutdown(shutdownCtx)
		if err != nil {
			s.logger.Fatal("Error shutting down", zap.Error(err))
		}
		s.logger.Info("Server shut down", zap.String("port", s.config.Address))
		serverStopCtx()
	}()

	s.logger.Info("Server started", zap.String("port", s.config.Address))
	err := serv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Fatal("Cant start server", zap.Error(err))
	}

	<-serverCtx.Done()
}
