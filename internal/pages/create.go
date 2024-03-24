package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type createRenderArgs struct {
	PreferredSuffix string
}

func CreatePage(c echo.Context) error {
	return c.Render(http.StatusOK, "create.html", createRenderArgs{})
}
