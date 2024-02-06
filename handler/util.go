package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component, layout templ.Component) error {
	return layout.Render(templ.WithChildren(c.Request().Context(), component), c.Response())
}
