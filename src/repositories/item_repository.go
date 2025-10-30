package repositories

import (
	"context"
	"spr-project/models"

	"gorm.io/gorm"
)

type ItemRepository struct {
	dB *gorm.DB
}

func (ItemRepository) New(db *gorm.DB) *ItemRepository {
	repo := ItemRepository{dB: db}
	return &repo
}

func (repo *OrderRepository) AddItem(ctx context.Context,
	orderId int64, productId int64, quantity int) (*models.Item, error) {
	order := models.Item{
		OrderId:   orderId,
		ProductId: productId,
		Quantity:  quantity,
	}

	result := repo.dB.WithContext(ctx).
		Create(&order)

	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}
