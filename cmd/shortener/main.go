package main

import (
	"github.com/dupreehkuda/ozon-shortener/internal/config"
	grpcHandlers "github.com/dupreehkuda/ozon-shortener/internal/handlers/grpc"
	httpHandlers "github.com/dupreehkuda/ozon-shortener/internal/handlers/http"
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
		db := postgres.New(cfg.DatabasePath, log)
		db.CreateSchema()
		store = db
	}

	serv := service.New(store, log)
	httpHandle := httpHandlers.New(serv, log)
	gRPCHandle := grpcHandlers.New(serv, log)

	srv := server.NewByConfig(httpHandle, gRPCHandle, log, cfg)
	srv.Run()
}
