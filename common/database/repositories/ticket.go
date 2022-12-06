package repositories

import (
	"gorm.io/gorm"
	"test_ticket/common/database/models"
)

type Persistor interface {
	Create(ticket *models.Ticket) error
}

type Ticket struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *Ticket {
	return &Ticket{db: db}
}

func (t *Ticket) Create(ticket *models.Ticket) error {
	return t.db.Create(ticket).Error
}
