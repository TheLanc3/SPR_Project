package terminal

import (
	"context"
	"fmt"
	"spr-project/repositories"
	"time"

	"gorm.io/gorm"
)

func Terminal(db *gorm.DB) bool {
	var choice int
	retVal := true

	for retVal {
		fmt.Println("1. To know the number of products in the stock inventory.")
		fmt.Println("2. To create a new order.")
		fmt.Println("3. To exit from application.")
		fmt.Println("Please enter number: 1, 2 or 3 on your keyboard.")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
		}

		switch choice {
		case 1:
			productRepo := repositories.NewProductRepository(db)
			suppRepo := repositories.NewSupplierRepository(db)
			// Create a context with a timeout
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			numberOfProducts, err := productRepo.NumberOfProducts(ctx)
			if err != nil {
				s := fmt.Errorf("func (repo *repositories.ProductRepository) NumberOfProducts return error: %s", err)
				fmt.Println(s)
			} else {
				fmt.Printf("The number of products in the stock inventory: %d\n", numberOfProducts)
			}
			details := true
			for details {
				fmt.Printf("Enter a number in the range from 1 to %d for details about particular product or anyother key to exit this section.\n",
					numberOfProducts)
				var iD int64
				_, err := fmt.Scanf("%d", &iD)
				if err != nil || iD < 1 || iD > numberOfProducts {
					details = false
					choice = 0
				} else {
					cTx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					defer cancel()
					prod, erroR := productRepo.GetProduct(cTx, iD)
					if erroR != nil {
						s := fmt.Errorf("func (repo *ProductRepository) GetProduct return error: %s", erroR)
						fmt.Println(s)
					} else {
						var suppName string
						ctX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
						defer cancel()
						supp, errors := suppRepo.GetSupplier(ctX, prod.SupplierId)
						if errors != nil {
							s := fmt.Errorf("func (repo *SupplierRepository) GetSupplier return error: %s", errors)
							fmt.Println(s)
						} else {
							suppName = supp.Name
						}
						fmt.Printf("Name:\t\t%s\n", prod.Name)
						fmt.Printf("Description:\t%s\n", prod.Description)
						fmt.Printf("Quantity:\t%d\n", prod.Quantity)
						fmt.Printf("Supplier:\t%s\n", suppName)
					}
				}
			}
		case 2:
			fmt.Println("Stub for an order creation.")
		case 3:
			retVal = false
		}
	}
	return retVal
}
