package ordersGenerator

import (
	"context"
	"fmt"
	"math/rand"
	"spr-project/parameters"
	"spr-project/repositories"
	"spr-project/services"
	"time"

	"gorm.io/gorm"
)

func Generator(doItAgain *bool, db *gorm.DB, customerId int64, step int64, pause time.Duration) (int, error) {
	rand.New(rand.NewSource(int64(time.Nanosecond)))
	productRepo := repositories.NewProductRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	totals := 0
	// Create a context with a timeout
	Ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	numberOfProducts, erR := productRepo.NumberOfProducts(Ctx)
	if erR != nil {
		return totals, erR
	} else {
		i := int64(rand.Intn(5)) + 1
		if i > numberOfProducts || i < 1 {
			i = 1
		}
		for i <= numberOfProducts {
			var numberOfItems int
			if 7 <= int(numberOfProducts) {
				numberOfItems = 7
			} else {
				numberOfItems = int(numberOfProducts)
			}
			numberOfItems = rand.Intn(numberOfItems) + 1
			orderServ := services.NewOrderService(db)
			var orderPositions []parameters.Position
			for k := 1; k <= numberOfItems; k++ {
				quantity := rand.Intn(5) + 1
				cTx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				prod, erroR := productRepo.GetProduct(cTx, i)
				if erroR != nil {
					return totals, erroR
				} else {
					orderPositions = append(orderPositions,
						parameters.Position{ProductId: i, Price: prod.Price,
							Quantity: quantity})
				}
				m := i + step
				if m > numberOfProducts {
					i = m - numberOfProducts
				} else {
					i = m
				}
			}
			if len(orderPositions) > 0 {
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				newOrder := parameters.NewOrderCreationData(customerId, orderPositions)
				order, err := orderServ.CreateNewOrder(ctx, newOrder)
				if err == nil {
					cust, eRr := customerRepo.GetCustomer(ctx, customerId)
					if eRr == nil {
						totals += order.Total
						fmt.Printf("New order with Id: %d, - was created for customer %s for amount %d\n",
							order.Id, cust.Name, order.Total)
					} else {
						return totals, eRr
					}
				} else {
					fmt.Errorf("error: %s, - to create an order", err)
				}
			}
			if !*doItAgain {
				return totals, nil
			} else {
				time.Sleep(pause * time.Second)
			}
		}
	}
	return totals, nil
}
