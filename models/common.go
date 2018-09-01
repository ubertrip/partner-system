package models

type WeeklyPayment struct {
	PaymentUuid   string  `json:"paymentUuid"`
	DriverUuid    string  `json:"driverUuid"`
	CashCollected float64 `json:"cashCollected"`
	Incentives    float64 `json:"incentives"`
	MiscPayment   float64 `json:"miscPayment"`
	NetFares      float64 `json:"netFares"`
	NetPayout     float64 `json:"netPayout"`
}

type Payment struct {
	DriverUuid    string  `json:"driverUuid"`
	PaymentUuid   string  `json:"paymentUuid"`
	StatementUuid string  `json:"statementUuid"`
	CreatedAt     string
	CreatedBy     string
	Credit        float64 `json:"credit"`
}

type Driver struct {
	Uuid string `json:"uuid"`
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
