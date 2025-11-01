package repositories

import (
	"context"
	"spr-project/enums"
	"spr-project/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	dB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	repo := OrderRepository{dB: db}
	return &repo
}

func (repo *OrderRepository) GetOrdersByCustomer(ctx context.Context,
	customerId int64, limit int) ([]models.Order, error) {
	var orders []models.Order

	if limit < 1 {
		limit = 3
	}

	result := repo.dB.WithContext(ctx).
		Preload("Positions").
		Where("customer_id = ?", customerId).
		Limit(limit).
		Find(&orders)

	if result.Error != nil {
		return []models.Order{}, result.Error
	}

	return orders, nil
}

func (repo *OrderRepository) GetUnfinishedOrdersByCustomer(ctx context.Context,
	customerId int64, limit int) ([]models.Order, error) {
	var orders []models.Order

	if limit < 1 {
		limit = 3
	}

	result := repo.dB.WithContext(ctx).
		Preload("Positions").
		Where("customer_id = ?", customerId).
		Where("status != ?", enums.Completed).
		Limit(limit).
		Find(&orders)

	if result.Error != nil {
		return []models.Order{}, result.Error
	}

	return orders, nil
}

func (repo *OrderRepository) GetDeliveredOrdersByCustomer(ctx context.Context,
	customerId int64, limit int) ([]models.Order, error) {
	var orders []models.Order

	if limit < 1 {
		limit = 3
	}

	result := repo.dB.WithContext(ctx).
		Preload("Positions").
		Where("customer_id = ?", customerId).
		Where("status != ?", enums.Delivered).
		Limit(limit).
		Find(&orders)

	if result.Error != nil {
		return []models.Order{}, result.Error
	}

	return orders, nil
}

func (repo *OrderRepository) AddOrder(ctx context.Context,
	customerId int64, total int) (*models.Order, error) {
	order := models.Order{
		CustomerId: customerId,
		Total:      total,
	}

	result := repo.dB.WithContext(ctx).
		Create(&order)

	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (repo *OrderRepository) UpdateStatus(ctx context.Context,
	newStatus enums.Status, id int64) error {
	result := repo.dB.WithContext(ctx).
		Model(&models.Order{}).
		Where("id = ?", id).
		Update("quantity", newStatus)

	return result.Error
}
