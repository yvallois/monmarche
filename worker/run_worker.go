package worker

import (
	"gorm.io/gorm/logger"
	"test_ticket/common"
	"test_ticket/common/database"
	"test_ticket/common/database/repositories"
	"test_ticket/worker/services"
)

func RunWorker(config common.Config) {
	database.Connect(
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
		10,
		logger.Info,
	)
	database.RunMigration()

	repository := repositories.NewTicketRepository(database.DB)
	createTicketService := services.CreateTicketService{Repository: repository}
	consumer := NewConsumer(config.RabbitMQURI, config.RabbitMQQueue, createTicketService.CreateTicket)
	defer consumer.Close()

	consumer.Run()
}
