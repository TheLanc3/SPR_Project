package repositories

import (
	"context"
	"spr-project/models"

	"gorm.io/gorm"
)

type StockRepository struct {
	dB *gorm.DB
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	repo := StockRepository{dB: db}
	return &repo
}

func (repo *StockRepository) GetStock(ctx context.Context,
	id int64) (models.Stock, error) {
	var stock models.Stock

	result := repo.dB.WithContext(ctx).
		Where("id = ?", id).
		First(&stock)

	if result.Error != nil {
		return models.Stock{}, result.Error
	}

	return stock, nil
}

func (repo *StockRepository) AddStock(ctx context.Context,
	name string, phone string, email string, address string) (*models.Stock, error) {
	stock := models.Stock{
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address}

	result := repo.dB.WithContext(ctx).
		Create(&stock)

	if result.Error != nil {
		return nil, result.Error
	}

	return &stock, nil
}
