package entity

import "time"

type Order struct {
	OrderUid          string `gorm:"primaryKey"`
	TrackNumber       string
	Entry             string
	Delivery          Delivery `gorm:"embedded;embeddedPrefix:delivery_"`
	Payment           Payment  `gorm:"embedded;embeddedPrefix:payment_"`
	Items             []Item   `gorm:"foreignKey:OrderRefer"`
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	Shardkey          string
	SmId              uint
	DateCreated       time.Time
	OofShard          string
}
