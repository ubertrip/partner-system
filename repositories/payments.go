package repositories

import (
	"github.com/ubertrip/partner-system/models"
	"strings"
	"fmt"
)

func UpdateWeekly(payments []models.WeeklyPayment) (ok bool) {
	values := make([]interface{}, 0, len(payments)*3)
	str := ""

	for _, payment := range payments {
		str += "(?, ?, ?),"
		values = append(values, payment.PaymentUuid, payment.DriverUuid, payment.CashCollected)
	}

	str = strings.TrimRight(str, ",")

	_, err := Get().Exec("INSERT INTO `weekly-payments` (paymentUuid, driverUuid, cashCollected) "+
		"VALUES "+ str + " ON DUPLICATE KEY UPDATE cashCollected=VALUES(cashCollected)", values...)

	if err != nil {
		fmt.Println(err)
		ok = false
		return
	}

	ok = true

	return
}
