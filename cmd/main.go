package main

import (
	"clean-reservations/config"
	db "clean-reservations/infrastructure/database"
	"clean-reservations/server"
	"context"
)

func main() {
	config.Init()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db.ConnectDB()

	server.Run(ctx)
}
