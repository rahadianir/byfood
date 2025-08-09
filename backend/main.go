package main

import (
	"byfood-app/internal/server"
	"context"
)

func main() {
	ctx := context.Background()
	server.StartServer(ctx)
}
