package verificator

import (
	"context"
	"fmt"
	"spr-project/parameters"
	"spr-project/repositories"
	"time"

	"gorm.io/gorm"
)

func Verifier(db *gorm.DB) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	productRepo := repositories.NewProductRepository(db)
	suppRepo := repositories.NewSupplierRepository(db)
	shipRepo := repositories.NewShipmentRepository(db)
	numberOfProducts, err := productRepo.NumberOfProducts(ctx)
	var i int64
	if err != nil {
		s := fmt.Errorf("func (repo *repositories.ProductRepository) NumberOfProducts return error: %s", err)
		fmt.Println(s)
	} else {
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
					var suppName string
					ctX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					defer cancel()
					shipmentExist, err := shipRepo.VerifyThatShipmenttAlreadyExist(ctX, prod.Id)
					if err != nil {
						s := fmt.Errorf("func (repo *ShipmentRepository) VerifyThatShipmenttAlreadyExist return error: %s", err)
						fmt.Println(s)
					} else {
						if !shipmentExist {
							supp, errors := suppRepo.GetSupplier(ctX, prod.SupplierId)
							if errors != nil {
								s := fmt.Errorf("func (repo *SupplierRepository) GetSupplier return error: %s", errors)
								fmt.Println(s)
							} else {
								suppName = supp.Name
							}
							fmt.Printf("%s product of %s supplier (with id: %d) has %d quantity.\n", prod.Name, prod.Supplier.Id, suppName, prod.Quantity)
							//Add new Shipment
							newShipment = append(newShipment, parameters.ShipmentData{prod.Id, 30, prod.Supplier.Id})
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
						}
					}
				}
			}
		}
	}
}
