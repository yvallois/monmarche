package helpers

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"test_ticket/common/database/models"
	"time"
)

func ProductParser(message string) (products []models.Product, err error) {
	csvReader := csv.NewReader(strings.NewReader(message))
	arrayBody, err := csvReader.ReadAll()
	if err != nil {
		return products, err
	}

	for _, body := range arrayBody[1:] {
		if len(body) != 3 {
			return products, fmt.Errorf("body should have three item per line")
		}
		price, err := strconv.ParseFloat(body[2], 32)
		if err != nil {
			return nil, fmt.Errorf("invalid price")
		}
		product := models.Product{
			ModelBase: models.ModelBase{
				UUID:      uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
			ProductID: body[1],
			Name:      body[0],
			Price:     float32(price),
		}
		products = append(products, product)
	}
	return products, err
}
