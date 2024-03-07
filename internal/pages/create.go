package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePage(c echo.Context) error {
	return c.Render(http.StatusOK, "create.html", "asd")
}
