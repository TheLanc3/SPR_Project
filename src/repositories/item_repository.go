package repositories

import (
	"context"
	"spr-project/models"
	"spr-project/parameters"

	"gorm.io/gorm"
)

type ItemRepository struct {
	dB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	repo := ItemRepository{dB: db}
	return &repo
}

func (repo *ItemRepository) AddItems(ctx context.Context,
	orderId int64, positions []parameters.Position) (*[]models.Item, error) {
	orderPositions := make([]models.Item, 0)
	for _, position := range positions {
		orderPositions = append(orderPositions, models.Item{
			OrderId:   orderId,
			ProductId: position.ProductId,
			Quantity:  position.Quantity,
		})
	}

	result := repo.dB.WithContext(ctx).
		Create(&orderPositions)

	if result.Error != nil {
		return nil, result.Error
	}

	return &orderPositions, nil
}
