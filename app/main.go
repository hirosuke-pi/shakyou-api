package main

import (
	"net/http"

	. "shakyou-api/api/v1"
	_ "shakyou-api/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Shakyou API
// @version 1.0
// @description S先生の写経課題を楽に終わらせるためのAPI

// @contact.name hirosuke-pi
// @contact.url https://www.twitter.com/hirosuke_pi

// @host localhost:8080
// @BasePath /api/v1

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	v1.GET("/", showHello)
	v1.POST("/shakyou", PostShakyouPdf)

	e.Logger.Fatal(e.Start(":8080"))
}

func showHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
