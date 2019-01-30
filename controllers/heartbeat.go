package controllers

import (
	"fmt"
	// "strings"

	// "log"
	"net/http"

	"time"

	"github.com/labstack/echo"

	"encoding/json"
	// "github.com\ubertrip\partner-system\controllers"
	"github.com/ubertrip/partner-system/repositories"
	"golang.org/x/crypto/bcrypt"
)
// middleware
func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// cookies, err := c.Cookie("sess")
		// fmt.Println(cookies.Name, "okey")
		// fmt.Println(cookies.Value, "okey")
		// if err != nil {
		// 	return JsonResponseErr(c, err)
		// }

		return next(c)
	}
}

func Info(c echo.Context) error {
	info := struct {
		Name string `json:"name"`
	}{"partner system v0.0"}
	return JsonResponseOk(c, info)
}

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Menu struct {
	Menu string `json:"Menu" `
}

type LoginStatus struct {
	Status bool `json:"status"`
}

type Cookie struct {
	Cookie string `json: "cookie"`
}

func Login(c echo.Context) error {
	var loginForm LoginForm

	//Cookies(r *http.Request, w http.ResponseWriter)

	json.NewDecoder(c.Request().Body).Decode(&loginForm)
	// readCookie(c echo.Context)

	var resp LoginStatus

	resp.Status = repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	ok := repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	if !ok {
		return JsonResponseErr(c, ok)
	}

	cookie := http.Cookie{
		Name:     "sess",
		Value:    "123AQW",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(48 * time.Hour),
	}
	// cookies, err := c.Cookie("sess")
	// if err != nil {
	// 	return err
	// }

	c.SetCookie(&cookie)

	// fmt.Println(cookies.Name)
	// fmt.Println(cookies.Value)

	return JsonResponseOk(c, resp)
}

const (
	defaultCost = 12
)

func init() {
	pwd := "123"
	hash := HashPassword(pwd)
	fmt.Println(pwd, hash)
	fmt.Println(CheckPassword(pwd, hash))
	// fmt.Println(CheckPassword(pwd, "$2a$12$yEm8VP3E8t6i6wzcA4AlgOK8olCf0/CTA0GiwYOc5B10Cn6htWA5K"))
}

func HashPassword(password string) string {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost); err == nil {
		return string(hashedPassword)
	}
	return ""
}

func CheckPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func Logout(c echo.Context) error {
	var loginForm LoginForm
	var cookie Cookie

	var resp LoginStatus

	resp.Status = repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	err := json.NewDecoder(c.Request().Body).Decode(&cookie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}
	// var resp LoginStatus

	cookies := http.Cookie{
		Name:   "sess",
		MaxAge: -1,
	}
	c.SetCookie(&cookies)
	fmt.Println("nice")
	return JsonResponseErr(c, cookies)

}
