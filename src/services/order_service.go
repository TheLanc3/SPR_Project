package services

import (
	"context"
	"fmt"
	"spr-project/enums"
	"spr-project/models"
	"spr-project/parameters"
	"spr-project/repositories"

	"gorm.io/gorm"
)

type OrderService struct {
	dB *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	service := OrderService{dB: db}
	return &service
}

func (service *OrderService) CreateNewOrder(ctx context.Context,
	data parameters.OrderCreationData) (*models.Order, error) {
	var result models.Order

	err := service.dB.Transaction(func(tx *gorm.DB) error {
		productRepo := repositories.NewProductRepository(tx)
		orderRepo := repositories.NewOrderRepository(tx)
		itemRepo := repositories.NewItemRepository(tx)

		for _, position := range data.Positions {
			if err := productRepo.DecreaseQuantity(ctx,
				position.Quantity, position.ProductId); err != nil {
				return err
			}
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

func (service *OrderService) UpdateOrderStatus(ctx context.Context,
	orderId int64, status enums.Status) error {
	err := service.dB.Transaction(func(tx *gorm.DB) error {
		orderRepo := repositories.NewOrderRepository(tx)
		if err := orderRepo.UpdateStatus(ctx, status, orderId); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (service *OrderService) GetOrdersByCustomer(ctx context.Context,
	customerId int64, filter enums.OrderType, limit int) (*[]models.Order, error) {
	var orders []models.Order
	var Error error
	if limit < 3 {
		limit = 5
	}

	repo := repositories.NewOrderRepository(service.dB)

	switch filter {
	case enums.All:
		{
			data, err := repo.GetOrdersByCustomer(ctx, customerId, limit)
			Error = err
			orders = *data
			break
		}
	case enums.OnlyUnfinished:
		{
			data, err := repo.GetUnfinishedOrdersByCustomer(ctx, customerId, limit)
			Error = err
			orders = *data
			break
		}
	case enums.OnlyDelivered:
		{
			data, err := repo.GetDeliveredOrdersByCustomer(ctx, customerId, limit)
			Error = err
			orders = *data
			break
		}
	}

	if Error != nil {
		return &[]models.Order{}, Error
	}

	return &orders, nil
}
