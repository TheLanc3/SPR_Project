package repositories

import (
	"context"
	"spr-project/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	dB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	repo := CustomerRepository{dB: db}
	return &repo
}

func (repo *CustomerRepository) GetCustomer(ctx context.Context,
	id int64) (models.Customer, error) {
	var customer models.Customer

	result := repo.dB.WithContext(ctx).
		Where("id = ?", id).
		First(&customer)

	if result.Error != nil {
		return models.Customer{}, result.Error
	}

	return customer, nil
}

func (repo *CustomerRepository) AddCustomer(ctx context.Context,
	name string, phone string, address string) (*models.Customer, error) {
	customer := models.Customer{
		Name:    name,
		Phone:   phone,
		Address: address}

	result := repo.dB.WithContext(ctx).
		Create(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (repo *CustomerRepository) NumberOfCustomers(ctx context.Context) (int64, error) {
	var count int64
	result := repo.dB.WithContext(ctx).
		Model(&models.Customer{}).
		Count(&count)

	if result.Error != nil {
		return -1, result.Error
	}

	return count, nil
}
