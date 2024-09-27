package model

type Products struct {
	ProductId          uint    `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductPrice       float32 `json:"product_price"`
	ProductRating      float32 `json:"product_rating"`
	ProductRateCount   int     `json:"product_rate_count"`
	ProductDescription string  `json:"product_description"`
	ProductImage       string  `json:"product_image"`
	CategoryId         int     `json:"category_id"`
	CategoryName       string  `json:"category_name"`
	IsFavorite         bool    `json:"is_favorite"`
	IsInCart           bool    `json:"is_in_cart"`
}
