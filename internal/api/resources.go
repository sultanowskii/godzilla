package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sultanowskii/godzilla/internal/util"
	"github.com/sultanowskii/godzilla/pkg/storage"
)

type (
	CreateResourceRequest struct {
		Url    string  `json:"url"`
		Suffix *string `json:"suffix"`
	}

	CreateResourceResponse struct {
		Url    string `json:"url"`
		Suffix string `json:"suffix"`
	}

	GetResourceResponse struct {
		Url    string `json:"url"`
		Suffix string `json:"suffix"`
	}
)

func CreateResource(c echo.Context) error {
	request := new(CreateResourceRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	urlString := request.Url
	if err := util.ValidateUrl(urlString); err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	client := storage.GetRedisClient()

	var suffix string
	if preferredSuffix := request.Suffix; preferredSuffix != nil {
		suffix = *preferredSuffix

		if err := util.ValidateCustomSuffix(suffix); err != nil {
			return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		}
	} else {
		suffix = util.GenerateSuffix(urlString)
	}

	// TODO: change to SetNX?
	client.Set(storage.Ctx, suffix, urlString, 0)

	response := &CreateResourceResponse{
		Url:    urlString,
		Suffix: suffix,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetResource(c echo.Context) error {
	suffix := c.Param("suffix")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, suffix).Result()

	if err != redis.Nil {
		return c.JSON(http.StatusNotFound, Error{Message: "not found"})
	}

	response := GetResourceResponse{
		Url:    url,
		Suffix: suffix,
	}

	return c.JSON(http.StatusOK, response)
}
