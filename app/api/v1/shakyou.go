package v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func postShakyouPdf(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func test1() {
	fmt.Println("aaaaaaaaaaaaaa")
}
