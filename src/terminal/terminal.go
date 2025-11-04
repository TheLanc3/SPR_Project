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
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		switch choice {
		case 1:
			productRepo := repositories.NewProductRepository(db)
			numberOfProducts, err := productRepo.NumberOfProducts(ctx)
			if err != nil {
				s := fmt.Errorf("func (repo *repositories.ProductRepository) NumberOfProducts return error: %s", err)
				fmt.Println(s)
			} else {
				fmt.Printf("The number of products in the stock inventory: %d\n", numberOfProducts)
			}

		case 2:
			fmt.Println("Stub for an order creation.")
		case 3:
			retVal = false
		}
	}
	return retVal
}
