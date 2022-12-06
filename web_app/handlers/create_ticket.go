package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"test_ticket/web_app/services"
)

type HandlerCreate struct {
	service services.ICreateTicketService
}

func NewHandlerCreateTicket(service services.ICreateTicketService) *HandlerCreate {
	return &HandlerCreate{service: service}
}

func (h *HandlerCreate) CreateTicketHandler(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.service.CreateTicketService(data)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
