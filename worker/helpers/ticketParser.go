package helpers

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
	"test_ticket/common/database/models"
	"time"
)

func TicketParser(message string) (*models.Ticket, error) {
	arrayMsg := strings.Split(strings.ReplaceAll(message, "\r", ""), "\n\n")
	if len(arrayMsg) != 2 {
		return nil, fmt.Errorf("message should have only one empty line")
	}
	ticket, err := validAndCreateTicket(arrayMsg[0])
	if err != nil {
		return nil, err
	}
	products, err := ProductParser(arrayMsg[1])
	if err != nil {
		return nil, err
	}
	ticket.Products = products
	return ticket, err
}

func validAndCreateTicket(header string) (*models.Ticket, error) {
	arrayHeader := strings.Split(header, "\n")
	if len(arrayHeader) != 3 {
		return nil, fmt.Errorf("invalid header")
	}

	r := regexp.MustCompile(`^Order: (\d+)$`)
	result := r.FindStringSubmatch(arrayHeader[0])
	if len(result) < 2 {
		return nil, fmt.Errorf("invalid orderId")
	}
	orderId, err := strconv.Atoi(result[1])
	if err != nil {
		return nil, fmt.Errorf("invalid orderId")
	}

	r = regexp.MustCompile(`^VAT: (\d+(\.\d+)?)$`)
	result = r.FindStringSubmatch(arrayHeader[1])
	if len(result) < 2 {
		return nil, fmt.Errorf("invalid vat")
	}
	vat, err := strconv.ParseFloat(result[1], 32)
	if err != nil {
		return nil, fmt.Errorf("invalid vat")
	}

	r = regexp.MustCompile(`^Total: (\d+(\.\d+)?)$`)
	result = r.FindStringSubmatch(arrayHeader[2])
	if len(result) < 2 {
		return nil, fmt.Errorf("invalid total")
	}
	total, err := strconv.ParseFloat(result[1], 32)
	if err != nil {
		return nil, fmt.Errorf("invalid total")
	}

	ticket := models.Ticket{
		ModelBase: models.ModelBase{
			UUID:      uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		OrderID: orderId,
		VAT:     float32(vat),
		Total:   float32(total),
	}
	return &ticket, nil
}
