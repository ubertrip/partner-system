package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	configuration "github.com/ubertrip/partner-system/config"
	"github.com/ubertrip/partner-system/controllers"
	"github.com/ubertrip/partner-system/repositories"
)

type JwtCustomClaims struct {
	Login bool `json:"login,omitempty"`
	ID    int  `json:"id,omitempty"`
	jwt.StandardClaims
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/JwtAuth")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderCookie,
			echo.HeaderSetCookie,
		},
		AllowCredentials: true,
	}))

	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &JwtCustomClaims{},
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:sess",
	}))

	repositories.InitDB()

	e.GET("/", controllers.Info)

	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.Login)
	e.GET("/statements", controllers.Accessible)

	e.GET("/logout", controllers.Logout)

	e.POST("/create-user", controllers.NewUser)

	e.POST("/payments", controllers.UpdateWeeklyPayments, controllers.Middleware)
	e.POST("/statements", controllers.UpdateWeeklyStatements, controllers.Middleware)
	e.PUT("/drivers/:id", controllers.UpdateDriver) // :driverUUID
	e.GET("/drivers/:id", controllers.GetDriver)    // :driverId

	e.GET("/statements", controllers.GetStatements, controllers.Middleware)

	e.POST("/credit/:uuid", controllers.AddCredit, controllers.Middleware)     // :driverUuuid
	e.GET("/credit/:uuid", controllers.GetByStatement, controllers.Middleware) // :statementUuid
	e.GET("/credit/:statementUUID/:driverUUID", controllers.GetDriverStatement, controllers.Middleware)

	e.Logger.Fatal(e.Start(":" + configuration.Get().Port))
}
