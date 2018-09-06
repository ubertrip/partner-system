package utils

import "github.com/ubertrip/partner-system/models"

func CalculateweeklyPaymentforDriver(weeklyPayments models.WeeklyPayment, payments []models.Payment) (report models.PaymentsReport) {
	report.Balance = 0

	for _, payment := range payments {
		report.Balance += payment.Credit
	}

	report.Diff = weeklyPayments.CashCollected-report.Balance

	return
}
