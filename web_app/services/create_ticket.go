package services

import (
	"test_ticket/common"
	"test_ticket/web_app/external"
)

type ICreateTicketService interface {
	CreateTicketService(body []byte) error
}

type CreateTicketService struct {
	publisher  external.Publisher
	exchange   string
	routingKey string
}

func NewCreateTicketService(publisher external.Publisher, config common.Config) *CreateTicketService {
	return &CreateTicketService{
		publisher:  publisher,
		exchange:   config.RabbitMQExchange,
		routingKey: config.RabbitMQRoutingKey,
	}
}

func (s *CreateTicketService) CreateTicketService(body []byte) error {
	err := s.publisher.Publish(s.exchange, s.routingKey, body)
	return err
}
