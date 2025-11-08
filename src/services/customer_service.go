package services

import (
	"context"
	"spr-project/models"
	"spr-project/parameters"
	"spr-project/repositories"

	"gorm.io/gorm"
)

type CustomerService struct {
	dB *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	service := CustomerService{dB: db}
	return &service
}

func (service *CustomerService) GetCustomerById(ctx context.Context,
	id int64) (*models.Customer, error) {
	repo := repositories.NewCustomerRepository(service.dB)

	customer, err := repo.GetCustomer(ctx, id)
	if err != nil {
		return &models.Customer{}, err
	}

	return &customer, nil
}

func (service *CustomerService) RegisterNewCustomer(ctx context.Context,
	data parameters.CustomerData) (*models.Customer, error) {
	var customer models.Customer

	err := service.dB.Transaction(func(tx *gorm.DB) error {
		repo := repositories.NewCustomerRepository(tx)

		result, err := repo.AddCustomer(ctx, data.Name, data.Phone,
			data.Address)
		if err != nil {
			return err
		}

		customer = *result

		return nil
	})

	if err != nil {
		return &models.Customer{}, err
	}

	return &customer, nil
}
