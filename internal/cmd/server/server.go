package server

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sultanowskii/godzilla/internal/api"
	"github.com/sultanowskii/godzilla/internal/pages"
)

func SetupEcho() *echo.Echo {
	t := &pages.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Pre(middleware.RemoveTrailingSlash())

	apiGroup := e.Group("/api")
	apiGroup.POST("/urls", api.CreateUrl)
	apiGroup.GET("/urls/:token", api.GetUrl)

	pageGroup := e.Group("")
	pageGroup.GET("/:token", pages.Dzilla)
	pageGroup.GET("/", pages.CreatePage)
	pageGroup.GET("/urls/:token", pages.UrlViewPage)

	e.Use(middleware.Logger())

	return e
}
