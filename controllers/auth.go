package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	configuration "github.com/ubertrip/partner-system/config"
	"github.com/ubertrip/partner-system/utils"

	m "github.com/ubertrip/partner-system/models"
	"github.com/ubertrip/partner-system/repositories"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	config = configuration.Get()

	tm = utils.Midnight().Add(config.CookieExpirationTime)

	claims = &m.JwtCustomClaims{
		"friend",
		true,
		jwt.StandardClaims{
			ExpiresAt: tm.Unix(),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ = token.SignedString([]byte("secret"))
)

// middleware
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if cookie, err := c.Cookie("sess"); err == nil {
			fmt.Println(cookie.Value)
			if cookie.Value == t {
				fmt.Println(cookie.Value, "value", t, "token")
				return next(c)
			}
		}
		return JsonResponseErr(c, "")
	}
}

type LoginForm struct {
	Login    string `json:"login" validate:"min=5,max=32,alphanum,required"`
	Password string `json:"password" validate:"min=5,max=32,alphanum,required"`
	ID       int    `json:"ID"`
}

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func Login(c echo.Context) error {

	var loginForm LoginForm

	json.NewDecoder(c.Request().Body).Decode(&loginForm)

	repositories.GetUserByLogin(loginForm.Login, loginForm.Password)

	cookie := http.Cookie{
		Name:     "sess",
		Value:    t,
		Path:     "/",
		HttpOnly: true,
		Expires:  tm,
	}
	c.SetCookie(&cookie)

	return JsonResponseOk(c, t)
}

func JWTAuth(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*m.JwtCustomClaims)
	name := claims.Name
	fmt.Println("+++++", user, claims, name)
	return JsonResponseOk(c, "Welcome "+name+"!")
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
	var validate *validator.Validate

	err := json.NewDecoder(c.Request().Body).Decode(&loginForm)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}

	validate = validator.New()

	user := validate.Struct(&loginForm)

	if loginForm.Login == "" || loginForm.Password == "" {
		return JsonResponseErr(c, "Please fill in the fields")
	}
	if user != nil {
		return JsonResponseErr(c, "Creating user use only numbers and letters, do not use less than four characters")
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
