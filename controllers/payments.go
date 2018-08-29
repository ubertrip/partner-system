package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	"github.com/ubertrip/partner-system/models"
	"github.com/ubertrip/partner-system/repositories"
)

func UpdateWeeklyPayments(c echo.Context) error {
	var weeklyPayments []models.WeeklyPayment

	err := json.NewDecoder(c.Request().Body).Decode(&weeklyPayments)

	if err != nil {
		return c.JSON(http.StatusOK, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}

	if repositories.UpdateWeekly(weeklyPayments) {
		return JsonResponseOk(c, weeklyPayments)
	}

	return JsonResponseErr(c, "Cannot update weekly payments")
}
