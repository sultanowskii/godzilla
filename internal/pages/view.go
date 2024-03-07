package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanowskii/godzilla/internal/api"
	"github.com/sultanowskii/godzilla/internal/models"
	"github.com/sultanowskii/godzilla/pkg/storage"
)

func UrlViewPage(c echo.Context) error {
	token := c.Param("token")

	client := storage.GetRedisClient()
	orig, err := client.Get(storage.Ctx, token).Result()

	url := &models.Url{
		Url:   orig,
		Token: token,
	}

	if err != nil {
		return c.JSON(http.StatusNotFound, api.Error{Message: "Not found."})
	}

	return c.Render(http.StatusOK, "dzilla.html", url)
}

func Dzilla(c echo.Context) error {
	token := c.Param("token")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, token).Result()

	if err != nil {
		return c.JSON(http.StatusNotFound, api.Error{Message: "Not found."})
	}

	return c.Redirect(http.StatusFound, url)
}
