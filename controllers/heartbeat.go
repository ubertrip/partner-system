package controllers

import (
	"github.com/labstack/echo"
	// "fmt"
	"net/http"
	"encoding/json"
	// "github.com\ubertrip\partner-system\controllers"
	"github.com/ubertrip/partner-system/repositories"
	
)

func Info(c echo.Context) error {
	info := struct{
		Name string `json:"name"`
	}{"partner system v0.0"}
	return JsonResponseOk(c, info)
}
	

type LoginForm struct {
	Login  string `json:"login" form:"login" query:"login"`
	Password string `json:"password" form:"password" query:"password"`
}

type LoginStatus struct {
	Status bool `json:"status"`
}
	
func Login(c echo.Context) error {
	var loginForm LoginForm

	json.NewDecoder(c.Request().Body).Decode(&loginForm)

	var resp LoginStatus
	resp.Status = repositories.GetUserByLogin(loginForm.Login, loginForm.Password)
	
	return c.JSON(http.StatusOK, resp)
}