package main

import (
	"byfood-app/internal/server"
	"context"
)

// @title ByFood App
// @version 1.0
// @description API documentation for ByFood App.
// @host localhost:8080
// @BasePath /
func main() {
	ctx := context.Background()
	server.StartServer(ctx)
}
