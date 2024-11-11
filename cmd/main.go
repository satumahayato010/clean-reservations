package main

import (
	"clean-reservations/config"
	"clean-reservations/server"
	"context"
)

func main() {
	config.Init()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.Run(ctx)
}
