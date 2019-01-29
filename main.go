package main

import (
	// "net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	configuration "github.com/ubertrip/partner-system/config"
	"github.com/ubertrip/partner-system/controllers"
	"github.com/ubertrip/partner-system/repositories"
	// "github.com/ubertrip/partner-system/utils"
	// "errors"
	// "fmt"
	// "time"
	// "fmt"
	// "net/http"
)

func main() {
	e := echo.New()

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

	// e.Use(func() echo.MiddlewareFunc {
	// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
	// 		return func(c echo.Context) error {
	// 			fmt.Println("good")
	// 			return next(c)
	// 		}
	// 	}
	// }())

	// e.Get(controllers.JwtAuth)

	repositories.InitDB()

	// e.GET("/", controllers.Info)

	// http.HandleFunc("/login", controllers.Middleware)
	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.Login)

	// e.POST("/login", utils.NewSessionStore)
	e.GET("/logout", controllers.Logout)

	e.POST("/payments", controllers.UpdateWeeklyPayments)
	e.POST("/statements", controllers.UpdateWeeklyStatements)
	e.PUT("/drivers/:id", controllers.UpdateDriver) // :driverUUID
	e.GET("/drivers/:id", controllers.GetDriver)    // :driverId

	e.GET("/statements", controllers.GetStatements)
	// , controllers.JwtAuth

	e.POST("/credit/:uuid", controllers.AddCredit)     // :driverUuuid
	e.GET("/credit/:uuid", controllers.GetByStatement) // :statementUuid
	e.GET("/credit/:statementUUID/:driverUUID", controllers.GetDriverStatement)

	e.Logger.Fatal(e.Start(":" + configuration.Get().Port))
}
