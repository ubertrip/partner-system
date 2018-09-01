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
	e.POST("/credit/:driverUUID", controllers.AddCredit)

	e.Logger.Fatal(e.Start(":"+configuration.Get().Port))
}
