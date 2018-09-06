package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	"github.com/ubertrip/partner-system/models"
	"github.com/ubertrip/partner-system/repositories"
	"github.com/satori/go.uuid"
	"github.com/ubertrip/partner-system/utils"
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

func GetDriverStatement(c echo.Context) error {
	statementUuid := c.Param("statementUUID")
	driverUuid := c.Param("driverUUID")

	driver, err := repositories.GetDriverByUUID(driverUuid)

	if err != nil {
		return JsonResponseErr(c, "Cannot select driver")
	}

	payments, err := repositories.GetDriverPaymetListByStatementId(statementUuid, driverUuid)

	//if err != nil {
	//	return JsonResponseErr(c, "Cannot select driver payments")
	//}

	weeklyPayment, err := repositories.GetDriverWeeklyPayment(statementUuid, driverUuid)

	//if err != nil {
	//	return JsonResponseErr(c, "Cannot select driver weekly payment")
	//}

	return JsonResponseOk(c, models.DriverSummary{
		driver,
		utils.CalculateweeklyPaymentforDriver(weeklyPayment, payments),
		weeklyPayment,
		payments})

}

func GetByStatement(c echo.Context) error {
	statementUuid := c.Param("uuid")

	sortedPayments := make(map[string][]models.Payment)
	sortedWeeklyPayments := make(map[string]models.WeeklyPayment)

	drivers, err := repositories.GetDriversByStatus("active")

	if err != nil {
		return JsonResponseErr(c, "Cannot select drivers")
	}

	payments, err := repositories.GetPaymentsByStatement(statementUuid)

	//if err != nil {
	//	return JsonResponseErr(c, "Cannot select payments")
	//}

	for _, payment := range payments {
		sortedPayments[payment.DriverUuid] = append(sortedPayments[payment.DriverUuid], payment)
	}

	weeklyPayments, err := repositories.GetWeeklyPaymentsByStatement(statementUuid)

	//if err != nil {
	//	return JsonResponseErr(c, "Cannot select weekly payments")
	//}

	for _, weeklyPayment := range weeklyPayments {
		sortedWeeklyPayments[weeklyPayment.DriverUuid] = weeklyPayment
	}

	driverSummares := make([]models.DriverSummary, len(drivers))

	for i, driver := range drivers {
		driverSummares[i].Driver = driver
		driverSummares[i].Report = utils.CalculateweeklyPaymentforDriver(sortedWeeklyPayments[driver.Uuid], sortedPayments[driver.Uuid])
		driverSummares[i].WeeklyPayment = sortedWeeklyPayments[driver.Uuid]
		driverSummares[i].Payments = sortedPayments[driver.Uuid]
	}

	return JsonResponseOk(c, driverSummares)

}

func GetStatements(c echo.Context) error {
	statements, err := repositories.GetStatements()

	if err != nil {
		return JsonResponseErr(c, "Cannot select statements")
	}

	return JsonResponseOk(c, statements)

}
