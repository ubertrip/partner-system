package controllers

import (
	"github.com/labstack/echo"
)

func Info(c echo.Context) error {
	info := struct{
		Name string `json:"name"`
	}{"partner system v0.0"}
	return JsonResponseOk(c, info)
}
