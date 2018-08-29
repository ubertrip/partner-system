package models

type WeeklyPayment struct {
	PaymentUuid   string  `json:"paymentUuid"`
	DriverUuid    string  `json:"driverUuid"`
	CashCollected float64 `json:"cashCollected"`
}

type Payment struct {
	DriverUuid  string
	PaymentUuid string `json:"paymentUuid"`
	CreatedAt   string
	Credit      float64
}
