package main

import (
	"callisplanics/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handler.LayoutHandler)

	e.Logger.Fatal(e.Start(":8085"))
}
