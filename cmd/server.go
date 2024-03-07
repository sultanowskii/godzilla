package main

import (
	"github.com/labstack/echo/v4"

	"github.com/sultanowskii/godzilla/internal/endpoints"
	"github.com/sultanowskii/godzilla/pkg/storage"
)

func main() {
	storage.InitRedisClient()

	e := echo.New()
	e.POST("/short", endpoints.CreateUrl)
	e.GET("/short", endpoints.GetUrl)

	e.Logger.Fatal(e.Start(":8431"))
}
