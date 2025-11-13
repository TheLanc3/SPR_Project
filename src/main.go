package main

import (
	"fmt"
	"log"
	"os"
	"spr-project/checkIn"
	"spr-project/database"
	"spr-project/loader"
	"spr-project/ordersGenerator"
	"spr-project/terminal"
	"spr-project/verificator"
	"time"

	"gorm.io/gorm"
)

func worker(id int, jobs <-chan int, results chan<- int, b *bool, db *gorm.DB) {
	retVal := 0
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(100 * time.Millisecond)
		switch j {
		case 1: //The terminal
			for *b {
				*b = terminal.Terminal(db)
			}
		case 2: //VerifiercheckIn
			for *b {
				verificator.Verifier(db)
				time.Sleep(100 * time.Second)
			}
		case 3: //CheckInDelivery
			for *b {
				checkIn.CheckInDelivery(db)
				time.Sleep(20 * time.Second)
			}
		case 4:
			total, err := ordersGenerator.Generator(b, db, 1, 2, 4)
			if err != nil {
				log.Fatal(err)
			}
			retVal = total
		case 5:
			total, err := ordersGenerator.Generator(b, db, 2, 3, 5)
			if err != nil {
				log.Fatal(err)
			}
			retVal = total
		case 6:
			total, err := ordersGenerator.Generator(b, db, 3, 5, 3)
			if err != nil {
				log.Fatal(err)
			}
			retVal = total
		case 7:
			total, err := ordersGenerator.Generator(b, db, 4, 7, 6)
			if err != nil {
				log.Fatal(err)
			}
			retVal = total
		case 8:
			total, err := ordersGenerator.Generator(b, db, 5, 11, 7)
			if err != nil {
				log.Fatal(err)
			}
			retVal = total
		case 9:
			total, err := ordersGenerator.Generator(b, db, 6, 13, 8)
			if err != nil {
				log.Fatal(err)
			}
			retVal = total
		case 10:
			total, err := ordersGenerator.Generator(b, db, 7, 6, 9)
			if err != nil {
				log.Fatal(err)
			}
			retVal = total
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
	dB := database.Init("stock.db")
	args := os.Args
	length := len(os.Args)
	uSage := "Usage: ./spr-project 'All Fresh.txt'"

	if length > 2 {
		fmt.Println(uSage)
		os.Exit(1)
	} else if length == 2 {
		filename := args[1]
		err := loader.ReadFile(filename, dB)
		if err != nil {
			log.Fatal(err)
		}
	}
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	work := true
	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results, &work, dB)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)
	totals := 0
	counter := 0
	for a := 1; a <= 10; a++ {
		value, ok := <-results
		if ok {
			totals += value
			if value > 0 {
				counter++
			}
		} else {
			fmt.Println("timeout has been received")
		}
	}
	if counter > 0 {
		if counter == 1 {
			fmt.Printf("%d customer bought goods on the %d rubles.\n", counter, totals)
		} else {
			fmt.Printf("%d customers bought goods on the %d rubles.\n", counter, totals)
		}
	}
}
