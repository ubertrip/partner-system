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

func UpdateWeeklyStatements(c echo.Context) error {
	var statements []models.Statement

	err := json.NewDecoder(c.Request().Body).Decode(&statements)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Error  string
		}{"error", err.Error()})
	}

	if repositories.UpdateStatements(statements) {
		return JsonResponseOk(c, statements)
	}

	return JsonResponseErr(c, "Cannot update statements")
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

	payment.DriverUuid = c.Param("uuid")
	payment.PaymentUuid = uuid.Must(uuid.NewV4()).String()

	if repositories.AddPayment(payment) {
		return JsonResponseOk(c, payment)
	}

	return JsonResponseErr(c, "Cannot add payment")
}

func GetDriverCredit(c echo.Context) error {
	statementUuid := c.Param("statementUUID")
	driverUuid := c.Param("driverUUID")

	creditReport, err := repositories.GetDriverCreditByStatement(statementUuid, driverUuid)

	if err != nil {
		return JsonResponseErr(c, "Cannot select payments sum")
	}

	payments, err := repositories.GetDriverPaymetListByStatementId(statementUuid, driverUuid)

	if err != nil {
		return JsonResponseErr(c, "Cannot select payments list")
	}

	return JsonResponseOk(c, struct {
		Report   models.CreditReport `json:"report"`
		Payments []models.Payment    `json:"payments"`
	}{creditReport, payments})

}

func GetByStatement(c echo.Context) error {
	statementUuid := c.Param("uuid")

	creditReports, err := repositories.GetCreditsByStatement(statementUuid)

	if err != nil {
		return JsonResponseErr(c, "Cannot select payments")
	}

	return JsonResponseOk(c, creditReports)

}

func GetStatements(c echo.Context) error {
	statements, err := repositories.GetStatements()

	if err != nil {
		return JsonResponseErr(c, "Cannot select statements")
	}

	return JsonResponseOk(c, statements)

}
