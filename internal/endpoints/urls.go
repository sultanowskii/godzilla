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
		Url  string `json:"url"`
		Hash string `json:"hash"`
	}

	GetUrlResponse struct {
		Url  string `json:"url"`
		Hash string `json:"hash"`
	}
)

func CreateUrl(c echo.Context) error {
	request := new(CreateUrlRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	url := request.Url
	hash := util.HashString(url)

	client := storage.GetRedisClient()
	client.Set(storage.Ctx, hash, url, 0)

	response := &CreateUrlResponse{
		Url:  url,
		Hash: hash,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetUrl(c echo.Context) error {
	hash := c.QueryParam("hash")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, hash).Result()

	if err != nil {
		return c.JSON(http.StatusNotFound, Error{Message: "Not found."})
	}

	response := GetUrlResponse{
		Url:  url,
		Hash: hash,
	}

	return c.JSON(http.StatusOK, response)
}
