package ticket_route

import (
	"github.com/gin-gonic/gin"
	"test_ticket/common"
	"test_ticket/web_app/external"
	"test_ticket/web_app/handlers"
	"test_ticket/web_app/services"
)

func InitTicketRouters(route *gin.Engine, publisher external.Publisher, config common.Config) {
	createTicketService := services.NewCreateTicketService(publisher, config)
	createTicketHandler := handlers.NewHandlerCreateTicket(createTicketService)

	route.GET("/health_check", handlers.HealthCheck)
	route.POST("/ticket", createTicketHandler.CreateTicketHandler)
}
