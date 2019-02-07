package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	configuration "github.com/ubertrip/partner-system/config"
	"github.com/ubertrip/partner-system/controllers"
	"github.com/ubertrip/partner-system/repositories"
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

	repositories.InitDB()

	e.GET("/", controllers.Info)

	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.Login)

	e.GET("/logout", controllers.Logout)

	e.POST("/payments", controllers.UpdateWeeklyPayments)
	e.POST("/statements", controllers.UpdateWeeklyStatements)
	e.PUT("/drivers/:id", controllers.UpdateDriver) // :driverUUID
	e.GET("/drivers/:id", controllers.GetDriver)    // :driverId

	e.GET("/statements", controllers.GetStatements, controllers.JwtAuth)

	e.POST("/credit/:uuid", controllers.AddCredit)     // :driverUuuid
	e.GET("/credit/:uuid", controllers.GetByStatement) // :statementUuid
	e.GET("/credit/:statementUUID/:driverUUID", controllers.GetDriverStatement)

	e.Logger.Fatal(e.Start(":" + configuration.Get().Port))
}
