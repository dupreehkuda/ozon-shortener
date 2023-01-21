package main

import (
	"github.com/dupreehkuda/ozon-shortener/internal/config"
	handlers "github.com/dupreehkuda/ozon-shortener/internal/handlers/http"
	i "github.com/dupreehkuda/ozon-shortener/internal/interfaces"
	"github.com/dupreehkuda/ozon-shortener/internal/logger"
	"github.com/dupreehkuda/ozon-shortener/internal/server"
	"github.com/dupreehkuda/ozon-shortener/internal/service"
	"github.com/dupreehkuda/ozon-shortener/internal/storage/memory"
	"github.com/dupreehkuda/ozon-shortener/internal/storage/postgres"
)

func main() {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	var store i.Storage
	if cfg.UseMemory {
		store = memory.New(log)
	} else {
		store = postgres.New(cfg.DatabasePath, log)
	}

	serv := service.New(store, log)
	handle := handlers.New(serv, log)

	srv := server.NewByConfig(handle, log, cfg)
	srv.Run()
}
