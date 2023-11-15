package main

import (
	"fmt"
	"log"
	"mongoApi/router"
	"net/http"
)

func init() {
	fmt.Println("Init Function Started")
}

func main() {
	fmt.Println("Mongodb API")

	router := router.Router()

	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("Service is Getting Started")
	fmt.Println("Server Listening at http://localhost:8080")
}
