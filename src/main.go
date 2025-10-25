package main

import (
	"fmt"
	"spr-project/database"
)

func main() {
	database.Init("hello.db")
	fmt.Println("Success")
}
