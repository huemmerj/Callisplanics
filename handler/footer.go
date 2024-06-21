package handler

import (
	"callisplanics/components"

	"github.com/labstack/echo/v4"
)

func FooterHandler(c echo.Context) error {
	return HTML(c, components.Footer())
}
