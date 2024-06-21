package main

import (
	"callisplanics/layout"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return HTML(c, layout.Layout())
	})
	component := hello("Jens es")

	e.GET("/hello", func(c echo.Context) error {
		return HTML(c, component)
	})

	e.Logger.Fatal(e.Start(":8085"))
}

func HTML(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
