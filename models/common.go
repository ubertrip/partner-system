package models

import jwt "github.com/dgrijalva/jwt-go"

type WeeklyPayment struct {
	PaymentUuid   string  `json:"paymentUuid"`
	DriverUuid    string  `json:"driverUuid"`
	CashCollected float64 `json:"cashCollected"` // наличка
	Incentives    float64 `json:"incentives"`    // бонусы
	MiscPayment   float64 `json:"miscPayment"`   // возврат средств
	NetFares      float64 `json:"netFares"`      // общая сумма
	NetPayout     float64 `json:"netPayout"`     // безнал
	UpdatedAt     string  `json:"updatedAt"`
}

type Payment struct {
	DriverUuid    string  `json:"driverUuid"`
	PaymentUuid   string  `json:"paymentUuid"`
	StatementUuid string  `json:"statementUuid"`
	CreatedAt     string  `json:"createdAt"`
	CreatedBy     string  `json:"createdBy"`
	Credit        float64 `json:"credit"`
	CashCollected float64 `json:"cashCollected"`
	Balance       float64 `json:"balance"`
	Extra         string  `json:"extra"`
	Gas           float64 `json:"gas"`
	Petrol        float64 `json:"petrol"`
}

type Driver struct {
	Uuid   string `json:"uuid"`
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Photo  string `json:"photo"`
}

type Statement struct {
	Uuid         string  `json:"uuid"`
	IsPaid       bool    `json:"isPaid"`
	CurrencyCode string  `json:"currencyCode"`
	StartAt      string  `json:"startAt"`
	EndAt        string  `json:"endAt"`
	Total        float64 `json:"total"`
	Timezone     string  `json:"timezone"`
	Hidden       bool    `json:"hidden"`
}

//type CreditReport struct {
//	Diff       float64 `json:"diff"`
//	DriverUuid string  `json:"driverUuid"`
//	DriverName string  `json:"driverName"`
//
//	Credit        float64 `json:"credit"`
//	CashCollected float64 `json:"cashCollected"`
//	Incentives    float64 `json:"incentives"`
//	MiscPayment   float64 `json:"miscPayment"`
//	NetFares      float64 `json:"netFares"`
//	NetPayout     float64 `json:"netPayout"`
//}

type PaymentsReport struct {
	Balance float64 `json:"balance"`
	Diff    float64 `json:"diff"`
	Gas     float64 `json:"gas"`
	Petrol  float64 `json:"petrol"`
}

type DriverSummary struct {
	Driver        Driver         `json:"driver"`
	Report        PaymentsReport `json:"report"`
	WeeklyPayment WeeklyPayment  `json:"weeklyPayment"`
	Payments      []Payment      `json:"payments"`
	Statement     Statement      `json:"statement"`
}

type JwtCustomClaims struct {
	Name  string `json:"name,omitempty"`
	Admin bool   `json:"admin,omitempty"`
	jwt.StandardClaims
}
