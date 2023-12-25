package main

import (
	"fmt"

	fetch "github.com/Galangrs/helper/fetch"
)

// ExampleResponse represents the response structure similar to a GORM model.
type ExampleResponse struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         string    `json:"created_at"`
	UpdatedAt         string    `json:"updated_at"`
	Products          []Product `json:"products"`
}

type Product struct {
	ID int `json:"id"`
}

func main() {
	headers := fetch.Header{
		"Content-Type":  "application/json",
		"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNzAzNjE0NDUyfQ.aVc4Go7YP2qvE2SM1kVyoxsea7UJV7L9pwqC4XlXbOY",
	}

	data := fetch.Data{}

	// Create an instance of ExampleResponse and pass its pointer to SendRequest.
	var exampleResponseInstance ExampleResponse
	response := fetch.SendRequest("POST", "http://localhost:8080/categories", data, headers, &exampleResponseInstance)
	if response.Err != nil {
		fmt.Println("success", response.StatusCode)
		fmt.Println("err", response.Err)
		return
	}

	fmt.Println("success", response.StatusCode)
}
