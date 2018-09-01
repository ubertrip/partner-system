package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	"github.com/ubertrip/partner-system/models"
	"github.com/ubertrip/partner-system/repositories"
	"github.com/satori/go.uuid"
)

func UpdateWeeklyPayments(c echo.Context) error {
	var weeklyPayments []models.WeeklyPayment

	err := json.NewDecoder(c.Request().Body).Decode(&weeklyPayments)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}

	if repositories.UpdateWeekly(weeklyPayments) {
		return JsonResponseOk(c, weeklyPayments)
	}

	return JsonResponseErr(c, "Cannot update weekly payments")
}

func AddCredit(c echo.Context) error {
	var payment models.Payment

	err := json.NewDecoder(c.Request().Body).Decode(&payment)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}

	payment.DriverUuid = c.Param("driverUUID")
	payment.PaymentUuid = uuid.Must(uuid.NewV4()).String()

	if repositories.AddPayment(payment) {
		return JsonResponseOk(c, payment)
	}

	return JsonResponseErr(c, "Cannot add payment")
}
