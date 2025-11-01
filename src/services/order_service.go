package services

import (
	"context"
	"fmt"
	"spr-project/models"
	"spr-project/parameters"
	"spr-project/repositories"

	"gorm.io/gorm"
)

type OrderService struct {
	dB *gorm.DB
}

func (service *OrderService) CreateNewOrder(ctx context.Context,
	data parameters.OrderCreationData) (*models.Order, error) {
	var result models.Order

	err := service.dB.Transaction(func(tx *gorm.DB) error {
		productRepo := repositories.NewProductRepository(tx)
		orderRepo := repositories.NewOrderRepository(tx)
		itemRepo := repositories.NewItemRepository(tx)

		for _, position := range data.Positions {
			productRepo.DecreaseQuantity(ctx, position.Quantity, position.ProductId)
		}

		order, err := orderRepo.AddOrder(ctx, data.CustomerId, data.Total)
		if err != nil {
			return err
		}
		result = *order

		if _, err := itemRepo.AddItems(ctx, order.Id, data.Positions); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create order transaction: %w", err)
	}

	return &result, nil
}
