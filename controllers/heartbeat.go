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
// func Login(c echo.Context) error {
// 	password := c.FormValue("password")	
// 	login := c.FormValue("login")
// 	return c.String(http.StatusOK, "login:" + login + ", password:" + password)}
	// fmt.Println("Login:", password, login)

	// return nil
	// return JsonResponseOk(c, Login)		

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



// func Login(r *http.Request, defaultCode int) (int, string) {
// 	p := string.Split(r.URL.Path, "/")
// 	if len(p) == 1 {
// 			return defaultCode, p[0]
// 	} else if len(p) > 1 {
// 			code, err := strconv.Atoi(p[0])
// 			if err == nil {
// 					return code, p[1]
// 			} else {
// 					return defaultCode, p[1]
// 			}
// 	} else {
// 			return defaultCode, ""
// 	}
// }
// if (isset($_GET["id"]))	
// $id = $_GET["id"];
// $c = $_GET['a'] + $_GET['b'];
// GET /path/resource?param1=value1&param2=value2 HTTP/1.1
	
	
	
