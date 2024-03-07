package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanowskii/godzilla/internal/util"
	"github.com/sultanowskii/godzilla/pkg/storage"
)

type (
	CreateUrlRequest struct {
		Url string `json:"url"`
	}

	CreateUrlResponse struct {
		Url    string `json:"url"`
		Dzilla string `json:"dzilla"`
	}

	GetUrlResponse struct {
		Url    string `json:"url"`
		Dzilla string `json:"dzilla"`
	}
)

func CreateUrl(c echo.Context) error {
	request := new(CreateUrlRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	url := request.Url
	dzilla := util.GetDzilla(url)

	client := storage.GetRedisClient()
	client.Set(storage.Ctx, dzilla, url, 0)

	response := &CreateUrlResponse{
		Url:    url,
		Dzilla: dzilla,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetUrl(c echo.Context) error {
	dzilla := c.QueryParam("dzilla")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, dzilla).Result()

	if err != nil {
		return c.JSON(http.StatusNotFound, Error{Message: "Not found."})
	}

	response := GetUrlResponse{
		Url:    url,
		Dzilla: dzilla,
	}

	return c.JSON(http.StatusOK, response)
}

func Dzilla(c echo.Context) error {
	dzilla := c.Param("dzilla")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, dzilla).Result()

	if err != nil {
		return c.JSON(http.StatusNotFound, Error{Message: "Not found."})
	}

	return c.Redirect(http.StatusFound, url)
}
