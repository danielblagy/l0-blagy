package entity

import "time"

type Order struct {
	OrderUid    string
	TrackNumber string
	Entry       string
	Delivery    Delivery `gorm:"embedded;embeddedPrefix:delivery_"`
	Payment     Payment  `gorm:"embedded;embeddedPrefix:payment_"`
	// items
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	Shardkey          string
	SmId              uint
	DateCreated       time.Time
	OofShard          string
}
