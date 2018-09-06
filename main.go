package main

import (
	"github.com/labstack/echo"
	"github.com/ubertrip/partner-system/controllers"
	"github.com/ubertrip/partner-system/repositories"
	"github.com/labstack/echo/middleware"
	configuration "github.com/ubertrip/partner-system/config"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods:  []string{"GET", "POSt", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	repositories.InitDB()

	e.GET("/", controllers.Info)
	e.POST("/payments", controllers.UpdateWeeklyPayments)
	e.POST("/statements", controllers.UpdateWeeklyStatements)
	e.PUT("/drivers/:id", controllers.UpdateDriver) // :driverUUID
	e.GET("/drivers/:id", controllers.GetDriver) // :driverId

	e.GET("/statements", controllers.GetStatements)

	e.POST("/credit/:uuid", controllers.AddCredit) // :driverUuuid
	e.GET("/credit/:uuid", controllers.GetByStatement) // :statementUuid
	e.GET("/credit/:statementUUID/:driverUUID", controllers.GetDriverStatement)

	e.Logger.Fatal(e.Start(":"+configuration.Get().Port))
}
