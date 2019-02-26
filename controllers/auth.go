package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	configuration "github.com/ubertrip/partner-system/config"

	"github.com/ubertrip/partner-system/repositories"
	"github.com/ubertrip/partner-system/utils"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
)

// middleware
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if cookie, err := c.Cookie("sess"); err == nil {
			fmt.Println(cookie.Value)
			return next(c)
		}
		return JsonResponseErr(c, "")
	}
}

type LoginForm struct {
	Login    string `json:"login" validate:"min=5,max=32,alphanum,required"`
	Password string `json:"password" validate:"min=5,max=32,alphanum,required"`
	ID       int    `json:"ID"`
}

type JwtCustomClaims struct {
	Login bool `json:"login"`
	ID    int  `json:"id"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {

	var loginForm LoginForm

	json.NewDecoder(c.Request().Body).Decode(&loginForm)

	login, id := repositories.GetUserByLogin(loginForm.Login, loginForm.Password)
	if !login {
		return JsonResponseErr(c, login)
	}

	config := configuration.Get()
	tm := utils.Midnight().Add(config.CookieExpirationTime)

	claims := &JwtCustomClaims{
		login,
		id,
		jwt.StandardClaims{
			ExpiresAt: tm.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("+")
		return err
	}

	cookie := http.Cookie{
		Name:     "sess",
		Value:    t,
		Path:     "/",
		HttpOnly: true,
		Expires:  tm,
	}

	c.SetCookie(&cookie)
	fmt.Println(cookie.Expires)
	fmt.Println(claims.ID)
	fmt.Println(claims.Login)

	return JsonResponseOk(c, "")

}

func Accessible(c echo.Context) error {
	return JsonResponseOk(c, "Accessible")
}

func JwtAuth(c echo.Context) error {
	user := c.Get("secret").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Login
	fmt.Println("+++++", user, claims, name)
	fmt.Println()
	return JsonResponseOk(c, name)
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
