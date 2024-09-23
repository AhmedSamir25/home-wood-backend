package model

type Cart struct {
	Id        uint `json:"id"`
	ProductId uint `json:"product_id"`
	UserId    uint `json:"user_id"`
	ProductQt int  `json:"product_qt"`
}
