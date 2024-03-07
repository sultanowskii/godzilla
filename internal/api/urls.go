package api

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
		Url   string `json:"url"`
		Token string `json:"token"`
	}

	GetUrlResponse struct {
		Url   string `json:"url"`
		Token string `json:"token"`
	}
)

func CreateUrl(c echo.Context) error {
	request := new(CreateUrlRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	url := request.Url
	token := util.GetToken(url)

	client := storage.GetRedisClient()
	client.Set(storage.Ctx, token, url, 0)

	response := &CreateUrlResponse{
		Url:   url,
		Token: token,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetUrl(c echo.Context) error {
	token := c.Param("token")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, token).Result()

	if err != nil {
		return c.JSON(http.StatusNotFound, Error{Message: "Not found."})
	}

	response := GetUrlResponse{
		Url:   url,
		Token: token,
	}

	return c.JSON(http.StatusOK, response)
}
