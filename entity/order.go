package entity

import "time"

type Order struct {
	OrderUid          string    `json:"order_uid" gorm:"primaryKey"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery" gorm:"embedded;embeddedPrefix:delivery_"`
	Payment           Payment   `json:"payment" gorm:"embedded;embeddedPrefix:payment_"`
	Items             []Item    `json:"items" gorm:"foreignKey:OrderRefer"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              uint      `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}
