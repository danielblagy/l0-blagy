package entity

type Payment struct {
	Transaction  string
	RequestId    string
	Currency     string
	Provider     string
	Amount       uint
	PaymentDt    uint
	Bank         string
	DeliveryCost uint
	GoodsTotal   uint
	CustomFee    uint
}
