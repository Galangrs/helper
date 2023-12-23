package main

import (
	"fmt"

	fetch "github.com/Galangrs/helper/fetch"
)

func main() {
	headers := fetch.Header{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	data := fetch.Data{
		"name": "John",
		"age":  30,
	}

	response, jsonStatus, err := fetch.SendRequest("GET", "https://example.com", data, headers)
	if err != nil {
		if jsonStatus {
			fmt.Println("err json", response)
			return
		} else {
			fmt.Println("err", err)
			return
		}
	}
	if jsonStatus {
		fmt.Println("sucess json", response)
		return
	} else {
		fmt.Println("sucess", response)
		return
	}
}
