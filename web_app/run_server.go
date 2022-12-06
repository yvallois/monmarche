package web_app

import (
	"github.com/gin-gonic/gin"
	"test_ticket/common"
	"test_ticket/web_app/external"
	ticketRoute "test_ticket/web_app/routes"
)

func RunServer(config common.Config) {
	r := gin.New()
	r.Use(gin.Recovery())

	publisher := external.NewPublisher(common.Conf.RabbitMQURI)
	defer publisher.Close()

	ticketRoute.InitTicketRouters(r, publisher, config)
	_ = r.Run(":8081")
}
