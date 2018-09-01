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

	_, err := Get().Exec("INSERT INTO `weekly-payments` (paymentUuid, driverUuid, cashCollected, incentives, miscPayment, netFares, netPayout) "+
		"VALUES "+ str+ " ON DUPLICATE KEY UPDATE cashCollected=VALUES(cashCollected) incentives=VALUES(incentives) miscPayment=VALUES(miscPayment) netFares=VALUES(netFares) netPayout=VALUES(netPayout)", values...)

	if err != nil {
		fmt.Println(err)
		ok = false
		return
	}

	ok = true

	return
}

func AddPayment(p models.Payment) (ok bool) {
	_, err := Get().Exec("INSERT INTO payments (uuid, driverUuid, createdBy, credit, statementUuid) VALUES (?, ?, 'system', ?, ?)",
		p.PaymentUuid, p.DriverUuid, p.Credit, p.StatementUuid)

	if err != nil {
		fmt.Println(err)
		ok = false
		return
	}

	ok = true

	return
}
