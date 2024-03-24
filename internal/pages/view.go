package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanowskii/godzilla/internal/api"
	"github.com/sultanowskii/godzilla/internal/models"
	"github.com/sultanowskii/godzilla/pkg/storage"
)

func ResourceInfoPage(c echo.Context) error {
	suffix := c.Param("suffix")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, suffix).Result()

	resource := &models.Resource{
		Url:    url,
		Suffix: suffix,
	}

	if err != nil {
		return c.JSON(http.StatusNotFound, api.Error{Message: "Not found."})
	}

	return c.Render(http.StatusOK, "dzilla.html", resource)
}

func Dzilla(c echo.Context) error {
	suffix := c.Param("suffix")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, suffix).Result()

	if err != nil {
		return c.Render(http.StatusOK, "create.html", createRenderArgs{PreferredSuffix: suffix})
	}

	return c.Redirect(http.StatusFound, url)
}
