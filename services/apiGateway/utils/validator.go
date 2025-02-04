package utils

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func SetCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}
