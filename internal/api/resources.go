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

	url := request.Url
	preferredSuffix := request.Suffix

	client := storage.GetRedisClient()

	var suffix string
	if preferredSuffix == nil {
		suffix = util.GenerateSuffix(url)
	} else {
		suffix = *preferredSuffix
		suffixExists := client.Exists(storage.Ctx, suffix).Val()

		if !util.IsSuffixValid(suffix) {
			return c.JSON(http.StatusBadRequest, Error{Message: "Suffix contains invalid characters."})
		}

		if suffixExists == 1 {
			return c.JSON(http.StatusBadRequest, Error{Message: "Suffix already exists."})
		}
	}

	// TODO: change to SetNX?
	client.Set(storage.Ctx, suffix, url, 0)

	response := &CreateResourceResponse{
		Url:    url,
		Suffix: suffix,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetResource(c echo.Context) error {
	suffix := c.Param("suffix")

	client := storage.GetRedisClient()
	url, err := client.Get(storage.Ctx, suffix).Result()

	if err != redis.Nil {
		return c.JSON(http.StatusNotFound, Error{Message: "Not found."})
	}

	response := GetResourceResponse{
		Url:    url,
		Suffix: suffix,
	}

	return c.JSON(http.StatusOK, response)
}
