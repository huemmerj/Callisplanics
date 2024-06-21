package main

import (
	"callisplanics/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handler.HomeHandler)

	e.GET("/session", handler.SessionHandler)

	e.Logger.Fatal(e.Start(":8085"))
}
