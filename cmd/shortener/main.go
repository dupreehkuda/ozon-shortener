package main

import "github.com/dupreehkuda/ozon-shortener/internal/server"

func main() {
	srv := server.NewByConfig()
	srv.Run()
}
