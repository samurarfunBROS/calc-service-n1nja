package main

import (
	"calc-service-n1nja/internal/calculator"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", calculator.CalculateHandler)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
