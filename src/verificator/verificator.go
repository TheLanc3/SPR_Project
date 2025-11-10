package verificator

import (
	"context"
	"fmt"
	"spr-project/enums"
	"spr-project/mailing"
	"spr-project/models"
	"spr-project/parameters"
	"spr-project/repositories"
	"spr-project/supplier"
	"time"

	"gorm.io/gorm"
)

func DeliveryCheckIn(db *gorm.DB, deliveredShipment *models.Shipment) error {
	productRepo := repositories.NewProductRepository(db)
	shipRepo := repositories.NewShipmentRepository(db)
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	error := shipRepo.UpdateShipmentStatus(ctx, deliveredShipment.Id, enums.Delivered)
	if error != nil {
		return error
	}
	cTx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	erroR := productRepo.IncreaseQuantity(cTx, deliveredShipment.Quantity, deliveredShipment.ProductId)
	if erroR != nil {
		return erroR
	}
	cTX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := shipRepo.UpdateShipmentStatus(cTX, deliveredShipment.Id, enums.Completed)
	if err != nil {
		return err
	}

	return nil
}

func Verifier(db *gorm.DB) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	productRepo := repositories.NewProductRepository(db)
	suppRepo := repositories.NewSupplierRepository(db)
	shipRepo := repositories.NewShipmentRepository(db)
	stockRepo := repositories.NewStockRepository(db)
	numberOfProducts, err := productRepo.NumberOfProducts(ctx)
	var i int64
	if err != nil {
		s := fmt.Errorf("func (repo *repositories.ProductRepository) NumberOfProducts return error: %s", err)
		fmt.Println(s)
	} else {
		tx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		stk, erroR := stockRepo.GetStock(tx, 1)
		if erroR != nil {
			s := fmt.Errorf("func (repo *StockRepository) GetStock return error for Id 1: %s", erroR)
			fmt.Println(s)
		}
		for i = 1; i <= numberOfProducts; i++ {
			cTx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			prod, erroR := productRepo.GetProduct(cTx, i)
			if erroR != nil {
				s := fmt.Errorf("func (repo *ProductRepository) GetProduct return error: %s", erroR)
				fmt.Println(s)
			} else {
				var newShipment []parameters.ShipmentData
				if prod.Quantity < 10 {
					ctX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					defer cancel()
					shipmentExist, err := shipRepo.VerifyThatShipmenttAlreadyExist(ctX, prod.Id)
					if err != nil {
						s := fmt.Errorf("func (repo *ShipmentRepository) VerifyThatShipmenttAlreadyExist return error: %s", err)
						fmt.Println(s)
					} else {
						if !shipmentExist {
							fmt.Printf("%s product of supplier (id: %d) has %d quantity.\n", prod.Name, prod.SupplierId, prod.Quantity)
							//Add new Shipment
							newShipment = append(newShipment, parameters.ShipmentData{prod.Id, 30, prod.SupplierId})
						}
					}
				}
				if len(newShipment) > 0 {
					CTX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					defer cancel()
					shipments, erR := shipRepo.AddShipments(CTX, newShipment)
					if erR != nil {
						s := fmt.Errorf("func (repo *SupplierRepository) GetSupplier return error: %s", erR)
						fmt.Println(s)
					} else {
						for _, value := range *shipments {
							fmt.Printf("New shipment was created for %d product_id with id: %d\n", value.ProductId, value.Id)
							Tx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
							defer cancel()
							prod, erroR := productRepo.GetProduct(Tx, i)
							if erroR != nil {
								s := fmt.Errorf("func (repo *ProductRepository) GetProduct return error: %s", erroR)
								fmt.Println(s)
							} else {
								var suppMail string
								ctX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
								defer cancel()
								supp, errors := suppRepo.GetSupplier(ctX, prod.SupplierId)
								if errors != nil {
									s := fmt.Errorf("func (repo *SupplierRepository) GetSupplier return error: %s", errors)
									fmt.Println(s)
								} else {
									suppMail = supp.Email
								}
								mailing.SendEmail(suppMail, prod.Name, value.Quantity, int(value.Id), &stk)
								time.Sleep(5 * time.Second)
								//should be in future in separate thread
								supplier.Supplier(stk.Email, prod.Name, value.Quantity, int(value.Id), &supp)
								//should be in future in separate thread too
								cTX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
								defer cancel()
								err := shipRepo.UpdateShipmentStatus(cTX, value.Id, enums.Shipped)
								if err != nil {
									s := fmt.Errorf("func (repo *ShipmentRepository) UpdateShipmentStatus return error: %s", err)
									fmt.Println(s)
								}
								//should be in future in separate thread too
								erR := DeliveryCheckIn(db, &value)
								if erR != nil {
									s := fmt.Errorf("func (repo *ShipmentRepository) UpdateShipmentStatus return error: %s", erR)
									fmt.Println(s)
								}
							}
						}
					}
				}
			}
		}
	}
}
