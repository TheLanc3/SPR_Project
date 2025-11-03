package main

import (
	"fmt"
	"log"
	"os"
	"spr-project/database"
	"spr-project/loader"
	"spr-project/models"
	"spr-project/terminal"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func worker(id int, jobs <-chan int, results chan<- int, b *bool, supplierInventory []loader.ProductInventory) {
	retVal := 0
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(100 * time.Millisecond)
		switch j {
		case 1: //The terminal
			for *b {
				*b = terminal.Terminal(supplierInventory)
			}
		default: //The stub for other cases
			for *b {
				time.Sleep(time.Second)
			}
		}
		fmt.Println("worker", id, "finished job", j)
		results <- retVal
	}
}

func main() {
	database.Init("stock.db")
	args := os.Args
	filename := args[1]
	length := len(os.Args)
	uSage := "Usage: ./spr-project 'All Fresh.txt'"
	var supplierInventory []loader.ProductInventory
	if length > 2 {
		fmt.Println(uSage)
		os.Exit(1)
	} else if length == 2 {
		charToFind := "."
		var supplierName string
		index := strings.Index(filename, charToFind)
		if index != -1 {
			supplierName = filename[:index]
		} else {
			supplierName = filename
		}
		db, err := gorm.Open(sqlite.Open("stock.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		var supplier models.Supplier
		res := db.Where("name = ?", supplierName).First(&supplier)
		var supplierId int64
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				fmt.Printf("Supplier %s not found.\n", supplierName)
			} else {
				fmt.Printf("Error retrieving supplier: %v\n", res.Error)
			}
		} else {
			supplierId = supplier.Id
		}

		supplierInventory, err := loader.ReadFile(filename, supplierInventory)
		if err != nil {
			log.Fatal(err)
		}
		for _, value := range supplierInventory {
			// Retrieve a product by its name
			var product models.Product
			result := db.Where("name = ?", value.Name).First(&product)
			//productRepo := repositories.New
			if result.Error != nil {
				if result.Error == gorm.ErrRecordNotFound {
					db.FirstOrCreate(&models.Product{}, models.Product{Name: value.Name, Description: value.Description,
						Price: value.Price, Quantity: value.Quantity, SupplierId: supplierId})
				} else {
					fmt.Printf("Error retrieving product: %v\n", result.Error)
				}
			} else {
				//quantity := product.Quantity
				//repositories.ProductRepository.UpdateQuantity(quantity+value.Quantity, product.Id)
			}
		}

	}
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	work := true
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results, &work, supplierInventory)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 10; a++ {
		<-results
	}

}
