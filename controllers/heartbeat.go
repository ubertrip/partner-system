package controllers

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"

	"github.com/labstack/echo"
	configuration "github.com/ubertrip/partner-system/config"
	"github.com/ubertrip/partner-system/repositories"
	"github.com/ubertrip/partner-system/utils"
	"golang.org/x/crypto/bcrypt"
)

// middleware
func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookies, err := c.Cookie("sess")
		if err != nil {
			return JsonResponseErr(c, err)
		}

		fmt.Println(cookies.Name, "okey")
		fmt.Println(cookies.Value, "okey")
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
	Cookie string `json:"cookie"`
}

func Login(c echo.Context) error {
	var loginForm LoginForm
	json.NewDecoder(c.Request().Body).Decode(&loginForm)

	var resp LoginStatus

	resp.Status = repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	ok := repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	if !ok {
		return JsonResponseErr(c, ok)
	}
	config := configuration.Get()

	cookie := http.Cookie{
		Name:     "sess",
		Value:    "123AQW",
		Path:     "/",
		HttpOnly: true,
		Expires:  utils.Midnight().Add(time.Duration(config.Cookie) * time.Hour),
	}

	c.SetCookie(&cookie)
	fmt.Println(cookie.Expires)
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
	var cookie Cookie

	json.NewDecoder(c.Request().Body).Decode(&cookie)

	cookiess, err := c.Cookie("sess")
	fmt.Println(cookiess.Name, "okey1")
	fmt.Println(cookiess.Value, "okey1")
	if err != nil {
		return JsonResponseErr(c, err)
	}

	cookies := http.Cookie{
		Name:   "sess",
		MaxAge: -1,
	}
	c.SetCookie(&cookies)
	fmt.Println("nice")
	return JsonResponseOk(c, nil)

}
