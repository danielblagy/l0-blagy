package entity

type Payment struct {
	Transaction  string `json:"transaction" validate:"required,len=19"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       uint   `json:"amount" validate:"required,numeric"`
	PaymentDt    uint   `json:"payment_dt" validate:"required,numeric"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost uint   `json:"delivery_cost" validate:"required,numeric"`
	GoodsTotal   uint   `json:"goods_total" validate:"required,numeric"`
	CustomFee    uint   `json:"custom_fee" validate:"numeric"`
}
