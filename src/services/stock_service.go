package services

import (
	"context"
	"spr-project/models"
	"spr-project/repositories"

	"gorm.io/gorm"
)

type StockService struct {
	dB *gorm.DB
}

func NewStockService(db *gorm.DB) *StockService {
	service := StockService{dB: db}
	return &service
}

func (service *StockService) GetStockById(ctx context.Context,
	id int64) (*models.Stock, error) {
	repo := repositories.NewStockRepository(service.dB)
	data, err := repo.GetStock(ctx, id)
	if err != nil {
		return &models.Stock{}, err
	}

	return &data, nil
}

func (service *StockService) AddNewStock(ctx context.Context,
	name string, phone string, email string, address string) (*models.Stock, error) {
	var stock models.Stock

	err := service.dB.Transaction(func(tx *gorm.DB) error {
		repo := repositories.NewStockRepository(tx)

		data, err := repo.AddStock(ctx, name, phone, email, address)
		if err != nil {
			return err
		}
		stock = *data
		return nil
	})

	if err != nil {
		return &models.Stock{}, err
	}

	return &stock, nil
}
