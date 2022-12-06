package models

type Ticket struct {
	ModelBase
	OrderID  int     `gorm:"not null" json:"order_id"`
	VAT      float32 `gorm:"not null" json:"vat"`
	Total    float32 `gorm:"not null" json:"total"`
	Products []Product
}
