package main

import (
	"github.com/sultanowskii/godzilla/internal/cmd/server"
	"github.com/sultanowskii/godzilla/pkg/storage"
)

func main() {
	storage.InitRedisClient()

	e := server.SetupEcho()

	e.Logger.Info(e.Start(":8431"))
}
