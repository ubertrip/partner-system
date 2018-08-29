package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func JsonResponseOk(c echo.Context, u interface{}) error {
	return c.JSON(http.StatusOK, struct {
		Status string      `json:"status"`
		Result interface{} `json:"result"`
	}{"ok", u})
}

func JsonResponseErr(c echo.Context, u interface{}) error {
	return c.JSON(http.StatusBadRequest, struct {
		Status  string      `json:"status"`
		Message interface{} `json:"message"`
	}{"error", u})
}
