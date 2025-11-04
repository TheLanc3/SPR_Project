package repositories

import (
	"context"
	"errors"
	"spr-project/models"
	"spr-project/parameters"

	"gorm.io/gorm"
)

type ProductRepository struct {
	dB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	repo := ProductRepository{dB: db}
	return &repo
}

func (repo *ProductRepository) GetProduct(ctx context.Context,
	id int64) (*models.Product, error) {
	var product models.Product

	result := repo.dB.WithContext(ctx).
		Where("id = ?", id).
		First(&product)

	if result.Error != nil {
		return &models.Product{}, result.Error
	}

	return &product, nil
}

func (repo *ProductRepository) AddProducts(ctx context.Context,
	data []parameters.Product) (*[]models.Product, error) {
	var prepare []models.Product
	for _, product := range data {
		prepare = append(prepare, models.Product{
			Name:        product.Name,
			Description: product.Description,
			SupplierId:  product.SupplierId,
			Price:       product.Price,
			Quantity:    product.Quantity,
		})
	}

	result := repo.dB.WithContext(ctx).
		Create(&prepare)

	if result.Error != nil {
		return nil, result.Error
	}

	return &prepare, nil
}

func (repo *ProductRepository) UpdateDescription(ctx context.Context,
	newDescription string, id int64) error {
	result := repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", id).
		Update("description", newDescription)

	return result.Error
}

func (repo *ProductRepository) UpdatePrice(ctx context.Context,
	newPrice int, id int64) error {
	result := repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", id).
		Update("price", newPrice)

	return result.Error
}

func (repo *ProductRepository) IncreaseQuantity(ctx context.Context,
	increment int, id int64) error {
	result := repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", id).
		Update("quantity", gorm.Expr("quantity + ?", increment))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insufficient stock or product not found")
	}
	return nil
}

func (repo *ProductRepository) DecreaseQuantity(ctx context.Context,
	decrement int, id int64) error {
	result := repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", id).
		Where("quantity >= ?", decrement).
		Update("quantity", gorm.Expr("quantity - ?", decrement))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insufficient stock or product not found")
	}
	return nil
}

func (repo *ProductRepository) NumberOfProducts(ctx context.Context) (int64, error) {
	var count int64
	result := repo.dB.WithContext(ctx).
		Model(&models.Product{}).
		Count(&count)

	if result.Error != nil {
		return -1, result.Error
	}

	return count, nil
}
