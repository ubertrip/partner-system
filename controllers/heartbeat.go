package controllers

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/labstack/echo"
	configuration "github.com/ubertrip/partner-system/config"
	"github.com/ubertrip/partner-system/jwt"
	"github.com/ubertrip/partner-system/repositories"
	"github.com/ubertrip/partner-system/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	inst = jwt.New("ePXXC2v2YCzZFW9yU9Pu2mBc3GgefkEVf5zWhAw9YcvFb8Na")
)

// middleware
func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if cookie, err := c.Cookie("sess"); err == nil {
			if _, success := inst.Decode(cookie.Value); !success {
				return JsonResponseErr(c, err)
			}
		}

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
	ID       int    `json:"ID"`
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

	resp.Status, _ = repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	ok, id := repositories.GetUserByLogin(loginForm.Login, loginForm.Password)
	if !ok {
		return JsonResponseErr(c, ok)
	}

	fmt.Println(loginForm.ID)
	config := configuration.Get()

	tm := utils.Midnight().Add(config.Cookie)

	session := jwt.Session{
		ID:      int64(id),
		Expired: tm.Unix(),
	}

	cookie := http.Cookie{
		Name:     "sess",
		Value:    inst.Encode(&session),
		Path:     "/",
		HttpOnly: true,
		Expires:  tm,
	}

	c.SetCookie(&cookie)
	fmt.Println(cookie.Expires)
	fmt.Println(session.ID)
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
