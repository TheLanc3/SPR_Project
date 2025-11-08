package services

import (
	"context"
	"spr-project/enums"
	"spr-project/models"
	"spr-project/parameters"
	"spr-project/repositories"

	"gorm.io/gorm"
)

type SupplierService struct {
	dB *gorm.DB
}

func NewSupplierService(db *gorm.DB) *SupplierService {
	service := SupplierService{dB: db}
	return &service
}

func (service *SupplierService) GetSupplierById(ctx context.Context,
	id int64) (*models.Supplier, error) {
	repo := repositories.NewSupplierRepository(service.dB)
	supplier, err := repo.GetSupplier(ctx, id)

	return &supplier, err
}

func (service *SupplierService) RegisterSupplier(ctx context.Context,
	data parameters.SupplierData) (*models.Supplier, error) {
	var supplier models.Supplier

	err := service.dB.Transaction(func(tx *gorm.DB) error {
		repo := repositories.NewSupplierRepository(tx)

		result, err := repo.AddSupplier(ctx, data.Name, data.Phone,
			data.Email, data.Address)

		if err != nil {
			return err
		}

		supplier = *result

		return nil
	})

	if err != nil {
		return &models.Supplier{}, err
	}

	return &supplier, nil
}

func (service *SupplierService) RegisterProductShipments(ctx context.Context,
	data []parameters.ShipmentData) (*[]models.Shipment, error) {
	var shipments []models.Shipment

	err := service.dB.Transaction(func(tx *gorm.DB) error {
		repoShipment := repositories.NewShipmentRepository(tx)

		output, err := repoShipment.AddShipments(ctx, data)
		if err != nil {
			return err
		}

		shipments = *output
		return nil
	})

	if err != nil {
		return &[]models.Shipment{}, err
	}

	return &shipments, nil
}

func (service *SupplierService) UpdateProductShipmentsStatus(ctx context.Context,
	data []parameters.ShipmentUpdateData) error {
	err := service.dB.Transaction(func(tx *gorm.DB) error {
		repoShipment := repositories.NewShipmentRepository(tx)
		repoProduct := repositories.NewProductRepository(tx)

		for _, shipment := range data {
			shipmentData, err := repoShipment.GetShipmentById(ctx, shipment.Id)
			if err != nil {
				return err
			}
			if shipment.Status == enums.Delivered {
				if err := repoProduct.IncreaseQuantity(ctx, shipmentData.Quantity,
					shipmentData.ProductId); err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}
