package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"spr-project/database"
	"strconv"
	"strings"
)

type ProductCatalog struct {
	Name        string
	Description string
	Price       int
	Quantity    int
}

func readFile(filename string, prodSlice []ProductCatalog) ([]ProductCatalog, error) {
	file, err := os.Open(filename)
	if err != nil {
		return prodSlice, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		charToFind := "\t"
		var item ProductCatalog
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
						return prodSlice, fmt.Errorf("price is absent for Name: %s", item.Name)
					}
				} else {
					return prodSlice, fmt.Errorf("description is absent for Name: %s", item.Name)
				}
			} else {
				return prodSlice, fmt.Errorf("description and Price are absent for Name: %s", item.Name)
			}
		}
		prodSlice = append(prodSlice, item)
	}

	if err := scanner.Err(); err != nil {
		return prodSlice, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return prodSlice, nil
}

func main() {
	database.Init("stock.db")
	args := os.Args
	filename := args[1]
	length := len(os.Args)
	uSage := "Usage: ./spr-project 'All Fresh.txt'"
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
		var supplierCatalog []ProductCatalog
		supplierCatalog, err := readFile(filename, supplierCatalog)
		if err != nil {
			log.Fatal(err)
		}
		numberOfProducts := len(supplierCatalog)
		fmt.Printf("%d lines was read for %s supplier\n", numberOfProducts, supplierName)
		fmt.Println("The last item in the Catalog:")
		fmt.Printf("Name:\t\t %s\n", supplierCatalog[numberOfProducts-1].Name)
		fmt.Printf("Description:\t %s\n", supplierCatalog[numberOfProducts-1].Description)
		fmt.Printf("Price:\t\t %d\n", supplierCatalog[numberOfProducts-1].Price)
		fmt.Printf("Quantity:\t %d\n", supplierCatalog[numberOfProducts-1].Quantity)
	}
}
