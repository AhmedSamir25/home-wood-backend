package model

import "time"

type Orders struct {
	Id               uint      `json:"id"`
	ProductId        uint      `json:"product_id"`
	UserId           uint      `json:"user_id"`
	ProductQt        int       `json:"product_qt"`
	OrderPrice       float64   `json:"order_price"`
	DeliveryLocation string    `json:"delivery_location"`
	OrderDate        time.Time `json:"order_date"`
	DeliveryDate     time.Time `json:"delivery_date"`
}
