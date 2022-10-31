package entity

import "time"

type Order struct {
	OrderUid          string    `json:"order_uid" gorm:"primaryKey" validate:"required;len=19"`
	TrackNumber       string    `json:"track_number" validate:"required;len=14"`
	Entry             string    `json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" gorm:"embedded;embeddedPrefix:delivery_" validate:"required"`
	Payment           Payment   `json:"payment" gorm:"embedded;embeddedPrefix:payment_" validate:"required"`
	Items             []Item    `json:"items" gorm:"foreignKey:OrderRefer" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmId              uint      `json:"sm_id" validate:"required;numeric"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}
