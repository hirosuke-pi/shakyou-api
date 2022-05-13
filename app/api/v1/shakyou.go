package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostShakyouPdf(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
