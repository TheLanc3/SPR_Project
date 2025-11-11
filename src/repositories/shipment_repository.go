package repositories

import (
	"context"
	"spr-project/enums"
	"spr-project/models"
	"spr-project/parameters"

	"gorm.io/gorm"
)

type ShipmentRepository struct {
	dB *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) *ShipmentRepository {
	repo := ShipmentRepository{dB: db}
	return &repo
}

func (repo *ShipmentRepository) GetShipmentById(
	ctx context.Context,
	id int64) (*models.Shipment, error) {
	var shipment models.Shipment

	result := repo.dB.WithContext(ctx).
		Where("id = ?", id).
		Find(&shipment)

	if result.Error != nil {
		return &models.Shipment{}, nil
	}
	return &shipment, nil
}

func (repo *ShipmentRepository) AddShipments(
	ctx context.Context,
	data []parameters.ShipmentData) (*[]models.Shipment, error) {
	var shipments []models.Shipment

	for _, shipment := range data {
		shipments = append(shipments, models.Shipment{
			ProductId:  shipment.ProductId,
			Quantity:   shipment.Quantity,
			SupplierId: shipment.SupplierId,
		})
	}

	result := repo.dB.WithContext(ctx).
		Create(&shipments)

	if result.Error != nil {
		return nil, result.Error
	}

	return &shipments, nil
}

func (repo *ShipmentRepository) UpdateShipmentStatus(ctx context.Context,
	shipmentId int64, status enums.Status) error {
	result := repo.dB.WithContext(ctx).
		Model(&models.Shipment{}).
		Where("id = ?", shipmentId).
		Updates(map[string]interface{}{
			"Status": status,
		})

	return result.Error
}

func (repo *ShipmentRepository) VerifyThatShipmenttAlreadyExist(ctx context.Context,
	productId int64) (bool, error) {
	var counter int64
	retVal := false
	result := repo.dB.WithContext(ctx).
		Model(&models.Shipment{}).
		Where("product_id = ?", productId).
		Where("status <> ?", enums.Completed).
		Where("status <> ?", enums.Canceled).
		Count(&counter)
	if result.Error != nil {
		return retVal, result.Error
	}
	if counter > 0 {
		retVal = true
	}
	return retVal, nil
}
