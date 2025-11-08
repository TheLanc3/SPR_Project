package services

import (
	"context"
	"spr-project/models"
	"spr-project/parameters"
	"spr-project/repositories"
	"strings"

	"gorm.io/gorm"
)

type ProductService struct {
	dB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	service := ProductService{dB: db}
	return &service
}

func (service *ProductService) GetProductById(ctx context.Context,
	id int64) (*models.Product, error) {
	repo := repositories.NewProductRepository(service.dB)
	data, err := repo.GetProduct(ctx, id)
	if err != nil {
		return &models.Product{}, err
	}

	return data, nil
}

func (service *ProductService) AddProducts(ctx context.Context,
	products []parameters.Product) (*[]models.Product, error) {
	var output []models.Product
	err := service.dB.Transaction(func(tx *gorm.DB) error {
		repo := repositories.NewProductRepository(tx)

		data, err := repo.AddProducts(ctx, products)
		if err != nil {
			return err
		}

		output = *data

		return nil
	})

	if err != nil {
		return &[]models.Product{}, err
	}

	return &output, nil
}

func (service *ProductService) UpdateProductsInfo(ctx context.Context,
	data []parameters.ProductUpdate) error {
	err := service.dB.Transaction(func(tx *gorm.DB) error {
		repo := repositories.NewProductRepository(tx)

		for _, info := range data {
			if info.Price != 0 {
				err := repo.UpdatePrice(ctx, info.Price, info.Id)
				if err != nil {
					return err
				}
			}
			if strings.TrimSpace(info.Description) != "" {
				err := repo.UpdateDescription(ctx, info.Description, info.Id)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}
