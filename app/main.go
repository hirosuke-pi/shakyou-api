package main

import (
	"net/http"

	_ "main/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

// @title Shakyou API
// @version 1.0
// @description S先生の写経課題を楽に終わらせるためのAPI

// @contact.name hirosuke-pi
// @contact.url https://www.twitter.com/hirosuke_pi

// @host localhost:8080
// @BasePath /api/v1

func main() {
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! !!! test")
}
