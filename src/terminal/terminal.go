package terminal

import (
	"fmt"
	"spr-project/loader"
)

func Terminal(supplierInventory []loader.ProductInventory) bool {
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
			numberOfProducts := len(supplierInventory)
			fmt.Printf("The number of products in the stock inventory: %d", numberOfProducts)
		case 2:
			fmt.Println("Stub for an order creation.")
		case 3:
			retVal = false
		}
	}
	return retVal
}
