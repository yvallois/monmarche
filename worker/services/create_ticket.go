package services

import (
	"log"
	"test_ticket/common/database/repositories"
	"test_ticket/worker/helpers"
)

type CreateTicketService struct {
	Repository *repositories.Ticket
}

func (s *CreateTicketService) CreateTicket(body string) bool {
	log.Printf("Received a message:\n%s", body)
	ticket, err := helpers.TicketParser(body)
	if err != nil {
		log.Printf("An error occured during ticket creation %v", err)
		return false
	}

	err = s.Repository.Create(ticket)
	if err != nil {
		log.Printf("An error occured during ticket creation %v", err)
		return false
	}
	return true
}
