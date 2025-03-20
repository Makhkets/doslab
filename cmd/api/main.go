package main

import (
	"doslab/internal/app"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fmt.Println("The server is running at ==> localhost:8000")
	app.Server()
}
