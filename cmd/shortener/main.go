package main

import (
	"github.com/dupreehkuda/ozon-shortener/internal/config"
	handlers "github.com/dupreehkuda/ozon-shortener/internal/handlers/http"
	"github.com/dupreehkuda/ozon-shortener/internal/logger"
	"github.com/dupreehkuda/ozon-shortener/internal/server"
	"github.com/dupreehkuda/ozon-shortener/internal/service"
	"github.com/dupreehkuda/ozon-shortener/internal/storage/postgres"
)

func main() {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	store := postgres.New(cfg.DatabasePath, log)
	serv := service.New(store, log)

	handle := handlers.New(serv, log)

	srv := server.NewByConfig(handle, log, cfg)
	srv.Run()
}
