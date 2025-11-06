package main

import (
	"fmt"
	"log"
	"os"
	"spr-project/database"
	"spr-project/loader"
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
		case 2: //Verifier
			for *b {
				verificator.Verifier(db)
				time.Sleep(100 * time.Second)
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
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results, &work, dB)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 10; a++ {
		<-results
	}

}
