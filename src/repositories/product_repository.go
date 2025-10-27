package repositories

import (
	"context"
	"spr-project/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	dB *gorm.DB
}

func (ProductRepository) InitRepo(db *gorm.DB) *ProductRepository {
	repo := ProductRepository{dB: db}
	return &repo
}

func (repo *ProductRepository) GetProduct(ctx context.Context,
	id int64) (models.Product, error) {
	var product models.Product

	result := repo.dB.WithContext(ctx).
		Where("id = ?", id).
		First(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) AddProduct(ctx context.Context,
	name string, description string, price int, quantity int) (*models.Product, error) {
	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity}

	result := repo.dB.WithContext(ctx).
		Create(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRepository) UpdateDescription(ctx context.Context,
	newDescription string, id int64) {
	repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", id).
		Update("description", newDescription)
}

func (repo *ProductRepository) UpdatePrice(ctx context.Context,
	newPrice int, id int64) {
	repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", id).
		Update("price", newPrice)
}

func (repo *ProductRepository) UpdateQuantity(ctx context.Context,
	newQuantity int, id int64) error {
	result := repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", id).
		Update("quantity", newQuantity)

	return result.Error
}
