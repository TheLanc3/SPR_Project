package terminal

import (
	"context"
	"fmt"
	"spr-project/parameters"
	"spr-project/repositories"
	"spr-project/services"
	"time"

	"gorm.io/gorm"
)

func createOrder(db *gorm.DB) error {
	// Create a context with a timeout
	customerRepo := repositories.NewCustomerRepository(db)
	ctX, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	numOfCustomers, erroR := customerRepo.NumberOfCustomers(ctX)
	if erroR != nil {
		return erroR
	} else {
		fmt.Printf("The number of customers in the database: %d\n", numOfCustomers)
		productRepo := repositories.NewProductRepository(db)
		chooseCustomer := true
		for chooseCustomer {
			fmt.Printf("Enter a number in the range from 1 to %d to choose the customer for which you place the order "+
				"or press any key to exit from order creation section: \n", numOfCustomers)
			var iD int64
			_, err := fmt.Scanf("%d", &iD)
			if err != nil || iD < 1 || iD > numOfCustomers {
				return erroR
			} else {
				// Create a context with a timeout
				Ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				numberOfProducts, erR := productRepo.NumberOfProducts(Ctx)
				if erR != nil {
					return erR
				} else {
					orderServ := services.NewOrderService(db)
					var orderPositions parameters.OrderCreationData
					orderPositions.CustomerId = iD
					fmt.Printf("Enter the product ID and quautity pairs to create positions of the order\n"+
						"the product ID must be in the range from 1 to %d, set it to zerro to finish the position creation.\n", numberOfProducts)
					var ID int64
					ID = 1
					var qnty int
					for ID > 0 && ID <= numberOfProducts {
						_, err := fmt.Scanf("%d %d", &ID, &qnty)
						if err != nil || ID < 1 || ID > numberOfProducts {
							if len(orderPositions.Positions) > 0 {
								ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
								defer cancel()
								order, err := orderServ.CreateNewOrder(ctx, orderPositions)
								if err == nil {
									cust, eRr := customerRepo.GetCustomer(ctx, iD)
									if eRr == nil {
										fmt.Printf("New order with Id: %d, - was created for customer %s",
											order.Id, cust.Name)
									} else {
										return eRr
									}
								}
								return err
							}
						} else {
							orderPositions.Positions = append(orderPositions.Positions,
								parameters.Position{ProductId: ID, Quantity: qnty})
						}
					}
				}
			}
		}
	}
	return nil
}

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
			err := createOrder(db)
			if err != nil {
				s := fmt.Errorf("Order creation error: %s", err)
				fmt.Println(s)
			}
		case 3:
			retVal = false
		}
	}
	return retVal
}
