package models

import "github.com/google/uuid"

type Product struct {
	ModelBase
	ProductID  string  `gorm:"not null;type:varchar(64)" json:"product_id"`
	Name       string  `gorm:"type:varchar(128)" json:"name"`
	Price      float32 `gorm:"not null;type:numeric" json:"price"`
	Ticket     Ticket  `gorm:"foreignKey:TicketUUID;constraint:OnDelete:CASCADE;"`
	TicketUUID uuid.UUID
}
