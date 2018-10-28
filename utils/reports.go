package utils

import "github.com/ubertrip/partner-system/models"

func CalculateweeklyPaymentforDriver(weeklyPayments models.WeeklyPayment, payments []models.Payment) (report models.PaymentsReport) {
	report.Balance = 0
	report.Gas = 0
	report.Petrol = 0

	for _, payment := range payments {
		report.Balance += payment.Credit
		report.Gas += payment.Gas
		report.Petrol += payment.Petrol
	}

	report.Diff = weeklyPayments.CashCollected-report.Balance

	return
}
