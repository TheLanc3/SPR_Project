package loader

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"spr-project/models"
	"spr-project/repositories"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ProductInventory struct {
	Name        string
	Description string
	Price       int
	Quantity    int
}

func ReadFile(filename string, db *gorm.DB) error {
	var prodSlice []ProductInventory
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		charToFind := "\t"
		var item ProductInventory
		index := strings.Index(s, charToFind)
		if index != -1 {
			item.Name = s[:index]
			if index+1 <= len(s) {
				stringAfter := s[index+1:]
				j := strings.Index(stringAfter, charToFind)
				if j != -1 {
					item.Description = stringAfter[:j]
					if j+1 <= len(stringAfter) {
						priceQuantity := stringAfter[j+1:]
						m := strings.Index(priceQuantity, charToFind)
						if m != -1 {
							price, erR := strconv.Atoi(priceQuantity[:m])
							if erR != nil {
								log.Fatalf("Error converting Price string to int: %v, for Name: %s", err, item.Name)
							}
							item.Price = price
							quantity, erR := strconv.Atoi(priceQuantity[m+1:])
							if erR != nil {
								log.Fatalf("Error converting Price string to int: %v, for Name: %s", err, item.Name)
							}
							item.Quantity = quantity
						} else {
							price, erR := strconv.Atoi(priceQuantity)
							if erR != nil {
								log.Fatalf("Error converting Price string to int: %v, for Name: %s", err, item.Name)
							}
							item.Price = price
						}

					} else {
						return fmt.Errorf("price is absent for Name: %s", item.Name)
					}
				} else {
					return fmt.Errorf("description is absent for Name: %s", item.Name)
				}
			} else {
				return fmt.Errorf("description and Price are absent for Name: %s", item.Name)
			}
		}
		prodSlice = append(prodSlice, item)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	charToFind := "."
	var supplierName string
	index := strings.Index(filename, charToFind)
	if index != -1 {
		supplierName = filename[:index]
	} else {
		supplierName = filename
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
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for _, value := range prodSlice {
		// Retrieve a product by its name
		var product models.Product
		result := db.Where("name = ?", value.Name).First(&product)
		productRepo := repositories.NewProductRepository(db)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				db.FirstOrCreate(&models.Product{}, models.Product{Name: value.Name, Description: value.Description,
					Price: value.Price, Quantity: value.Quantity, SupplierId: supplierId})
			} else {
				fmt.Printf("Error retrieving product: %v\n", result.Error)
			}
		} else {
			err := productRepo.IncreaseQuantity(ctx, value.Quantity, product.Id)
			if err != nil {
				s := fmt.Errorf("error: %s, - to increase quantity for product with Id: %d", err, product.Id)
				fmt.Println(s)
			}
		}
	}

	return nil
}
