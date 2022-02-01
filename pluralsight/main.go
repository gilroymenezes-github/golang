package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {

	u := models.User {
		ID: 2, 
		FirstName: "Tricia",
		LastName: "McMillan",
	}
	fmt.Println(u)
	fmt.Println("Hello from module gophers!")
	port := 3000
	_, err := startWebServer(port, 3)
	fmt.Println(err);
}

func startWebServer(port, numberOfRetries int) (int, error) {

	fmt.Println("Starting the server...")

	fmt.Println("Server has startedon port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, nil
}