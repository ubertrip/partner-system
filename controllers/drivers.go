package controllers

import (
	"github.com/labstack/echo"
	"github.com/ubertrip/partner-system/models"
	"encoding/json"
	"net/http"
	"github.com/ubertrip/partner-system/repositories"
)

func UpdateDriver(c echo.Context) error {
	var driver models.Driver

	err := json.NewDecoder(c.Request().Body).Decode(&driver)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}

	driver.Uuid = c.Param("id")

	if repositories.UpdateDriver(driver) {
		return JsonResponseOk(c, driver)
	}

	return JsonResponseErr(c, "Cannot add driver")
}

func GetDriver(c echo.Context) error {
	driverID := c.Param("id")

	driver, err := repositories.GetDriverById(driverID)

	if err != nil {
		return JsonResponseErr(c, "Cannot found driver")
	}

	return JsonResponseOk(c, driver)
}