package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

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

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	ID       int    `json:"ID"`
}

type User struct {
	Login    string
	Password string
}

type Menu struct {
	Menu string `json:"Menu" `
}

type LoginStatus struct {
	Status bool `json:"status"`
}

func Login(c echo.Context) error {
	var loginForm LoginForm
	json.NewDecoder(c.Request().Body).Decode(&loginForm)

	var resp LoginStatus

	resp.Status, _ = repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	login, id := repositories.GetUserByLogin(loginForm.Login, loginForm.Password)
	if !login {
		return JsonResponseErr(c, login)
	}

	config := configuration.Get()
	tm := utils.Midnight().Add(config.CookieExpirationTime)

	session := jwt.Session{
		ID:      int64(id),
		Login:   login,
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
	// fmt.Println(session.Login)
	return JsonResponseOk(c, resp)
}

const (
	defaultCost = 12
)

func HashPassword(password string) string {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost); err == nil {
		return string(hashedPassword)
	}
	return ""
}

func CheckPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// curl -X POST http://localhost:4321/create-user -H "Content-Type: application/json" -d "{\"login\": \"test\", \"password\": \"123123\"}"
func NewUser(c echo.Context) error {
	var loginForm LoginForm

	err := json.NewDecoder(c.Request().Body).Decode(&loginForm)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	re.MatchString(loginForm.Login)
	re.MatchString(loginForm.Password)

	if loginForm.Login == "" || loginForm.Password == "" {
		return JsonResponseErr(c, "Please fill in the fields")
	}

	if re.MatchString(loginForm.Login) == false || re.MatchString(loginForm.Password) == false {
		return JsonResponseErr(c, "Creating user use only numbers and letters")
	}

	if err := repositories.CreateUser(loginForm.Login, HashPassword(loginForm.Password)); err != nil {
		return JsonResponseErr(c, err.Error())
	}

	return JsonResponseOk(c, "User was created")
}

func Logout(c echo.Context) error {

	_, err := c.Cookie("sess")
	if err != nil {
		return JsonResponseErr(c, err)
	}

	cookies := http.Cookie{
		Name:   "sess",
		MaxAge: -1,
	}
	c.SetCookie(&cookies)
	fmt.Println("nice you logout")
	return JsonResponseOk(c, nil)

}
