package repositories

import (
	"context"
	"spr-project/enums"
	"spr-project/models"

	"gorm.io/gorm"
)

type ShipmentRepository struct {
	dB *gorm.DB
}

func (ShipmentRepository) New(db *gorm.DB) *ShipmentRepository {
	repo := ShipmentRepository{dB: db}
	return &repo
}

func (repo *ShipmentRepository) GetShipment(ctx context.Context,
	id int64) (models.Shipment, error) {
	var shipment models.Shipment

	result := repo.dB.WithContext(ctx).
		Where("id = ?", id).
		First(&shipment)

	if result.Error != nil {
		return models.Shipment{}, result.Error
	}

	return shipment, nil
}

func (repo *ShipmentRepository) GetShipmentsByProduct(ctx context.Context,
	productId int64) ([]models.Shipment, error) {
	var shipments []models.Shipment

	result := repo.dB.WithContext(ctx).
		Where("product_id = ?", productId).
		Order("id DESC").
		Find(&shipments)

	if result.Error != nil {
		return []models.Shipment{}, result.Error
	}

	return shipments, nil
}

func (repo *ShipmentRepository) GetShipmentsBySupplier(ctx context.Context,
	supplierId int64) ([]models.Shipment, error) {
	var shipments []models.Shipment

	result := repo.dB.WithContext(ctx).
		Where("supplier_id = ?", supplierId).
		Order("id DESC").
		Find(&shipments)

	if result.Error != nil {
		return []models.Shipment{}, result.Error
	}

	return shipments, nil
}

func (repo *ShipmentRepository) AddShipment(
	ctx context.Context, productId int64, quantity int,
	supplierId int64, status enums.Status) (*models.Shipment, error) {
	stock := models.Shipment{
		ProductId:  productId,
		Quantity:   quantity,
		SupplierId: supplierId,
		Status:     status}

	result := repo.dB.WithContext(ctx).
		Create(&stock)

	if result.Error != nil {
		return nil, result.Error
	}

	return &stock, nil
}
