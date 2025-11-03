package loader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ProductInventory struct {
	Name        string
	Description string
	Price       int
	Quantity    int
}

func ReadFile(filename string, prodSlice []ProductInventory) ([]ProductInventory, error) {
	file, err := os.Open(filename)
	if err != nil {
		return prodSlice, fmt.Errorf("ошибка открытия файла: %w", err)
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
