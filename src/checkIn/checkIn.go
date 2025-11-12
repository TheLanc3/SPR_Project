package checkIn

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"spr-project/models"
	"spr-project/repositories"
	"spr-project/verificator"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func CheckInDelivery(db *gorm.DB) {
	directoryPath := "../data/"
	destinationDir := "../data/completed"
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	// Ensure the destination directory exists
	err = os.MkdirAll(destinationDir, 0755) // Creates the directory and any necessary parent directories
	if err != nil {
		fmt.Printf("Error creating destination directory: %v\n", err)
		return
	}
	// fmt.Printf("Files in directory '%s':\n", directoryPath)
	// Iterate through the directory entries
	for _, entry := range entries {
		// Check if the entry is a file (not a directory)
		if !entry.IsDir() {
			fileName := entry.Name()
			var supplierName string
			charToFind := "."
			index := strings.Index(fileName, charToFind)
			if index != -1 {
				supplierName = fileName[:index]
			} else {
				supplierName = fileName
			}
			now := time.Now()
			formattedTime := now.Format("20060102150405")
			stampedFile := supplierName + charToFind + formattedTime
			sourceFile := filepath.Join(directoryPath, fileName)
			destinationFile := filepath.Join(destinationDir, stampedFile)
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

			file, err := os.Open(sourceFile)
			if err != nil {
				fmt.Printf("ошибка открытия файла: %w", err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			var productName string
			var productQuantity int
			var shipmentId int64
			shipRepo := repositories.NewShipmentRepository(db)
			for scanner.Scan() {
				s := scanner.Text()
				charToFind := "\t"
				index := strings.Index(s, charToFind)
				if index != -1 {
					productName = s[:index]
					if index+1 <= len(s) {
						stringAfter := s[index+1:]
						j := strings.Index(stringAfter, charToFind)
						if j != -1 {
							quantity, erR := strconv.Atoi(stringAfter[:j])
							if erR != nil {
								log.Fatalf("Error converting Quantity string to int: %v, for Name: %s", err, productName)
							}
							productQuantity = quantity
							if j+1 <= len(stringAfter) {
								shippID := stringAfter[j+1:]
								m := strings.Index(shippID, charToFind)
								if m != -1 {
									shipmentID, erR := strconv.Atoi(shippID[:m])
									if erR != nil {
										log.Fatalf("Error converting Price string to int: %v, for Name: %s", err, productName)
									}
									shipmentId = int64(shipmentID)
								} else {
									shipmentID, erR := strconv.Atoi(shippID)
									if erR != nil {
										log.Fatalf("Error converting Price string to int: %v, for Name: %s", err, productName)
									}
									shipmentId = int64(shipmentID)
								}

							} else {
								fmt.Printf("supplier ID is absent for Name: %s", productName)
								return
							}
						} else {
							fmt.Printf("quantity is absent for Name: %s", productName)
							return
						}
					} else {
						fmt.Printf("quantity and supplier ID are absent for Name: %s", productName)
						return
					}
				}
				var product models.Product
				res := db.Where("name = ?", productName).First(&product)
				var productId int64
				if res.Error != nil {
					if res.Error == gorm.ErrRecordNotFound {
						fmt.Printf("Product name %s not found.\n", productName)
					} else {
						fmt.Printf("Error retrieving Product: %v\n", res.Error)
					}
				} else {
					productId = product.Id
				}
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				shipment, err := shipRepo.GetShipmentById(ctx, shipmentId)
				if err != nil {
					s := fmt.Errorf("func (repo *ShipmentRepository) GetShipmentById return error: %s", err)
					fmt.Println(s)
				} else {
					if supplierId == shipment.SupplierId {
						if productId == shipment.ProductId {
							if productQuantity == shipment.Quantity {
								erR := verificator.DeliveryCheckIn(db, shipment)
								if erR != nil {
									s := fmt.Errorf("func DeliveryCheckIn return error: %s", erR)
									fmt.Println(s)
								}
							} else {
								fmt.Printf("The product quantity[%d] for %s doesnot match the Quantity[%d] in the shipments table.", productQuantity,
									productName, shipment.Quantity)
							}
						} else {
							fmt.Printf("The product Id[%d] for %s doesnot match the ProductId[%d] in the shipments table.", productId,
								productName, shipment.ProductId)
						}
					} else {
						fmt.Printf("The supplier Id[%d] for %s doesnot match the SupplierId[%d] in the shipments table.", supplierId,
							supplierName, shipment.SupplierId)
					}
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("ошибка чтения файла: %w", err)
				return
			}

			err = os.Rename(sourceFile, destinationFile)
			if err != nil {
				fmt.Printf("Error moving file: %v\n", err)
				return
			}
		}
	}
}
