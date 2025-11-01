package repositories

import (
	"context"
	"spr-project/models"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	dB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	repo := SupplierRepository{dB: db}
	return &repo
}

func (repo *SupplierRepository) GetSupplier(ctx context.Context,
	id int64) (models.Supplier, error) {
	var supplier models.Supplier

	result := repo.dB.WithContext(ctx).
		Where("id = ?", id).
		First(&supplier)

	if result.Error != nil {
		return models.Supplier{}, result.Error
	}

	return supplier, nil
}

func (repo *SupplierRepository) AddSupplier(ctx context.Context,
	name string, phone string, email string, address string) (*models.Supplier, error) {
	supplier := models.Supplier{
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address}

	result := repo.dB.WithContext(ctx).
		Create(&supplier)

	if result.Error != nil {
		return nil, result.Error
	}

	return &supplier, nil
}
