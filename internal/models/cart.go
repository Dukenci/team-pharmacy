package models

type Cart struct {
	UserID     uint
	Items      []CartItem
	TotalPrice float64
}

type CartItem struct {
	ID           uint
	UserID       uint    `json:"-"`
	MedicineID   uint    `json:"medicine_id"`
	MedicineName string  `json:"medicine_name"`
	Quantity     int     `json:"quantity"`
	PricePerUnit int     `json:"price_per_unit"`
	LineTotal    float64 `json:"line_total"`
}

type CartUpsert struct {
	MedicineID uint `json:"medicine_id"`
	Quantity   int  `json:"quantity"`
}

type CartUpdateQuantity struct {
	Quantity int `json:"quantity"`
}