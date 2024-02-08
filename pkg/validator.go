package pkg

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func FormValidate(c echo.Context, m interface{}) error {
	if err := c.Bind(m); err != nil {
		return err
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(m)
	if err != nil {
		e := err.(validator.ValidationErrors)
		return e
	}
	return nil
}
