package fetch

import (
	"bytes"         // Used to create a buffer for the request body.
	"encoding/json" // Used to marshal the data to JSON.
	"errors"        // Used to create an error if the response status code is >= 400.
	"io/ioutil"     // Used to read the response body.
	"net/http"      // Used to send the HTTP request.
)

// Header struct represents the headers for the request
type Header map[string]string

// Data struct represents the data for the request
type Data map[string]interface{}

/*
SendRequest sends an HTTP request and returns the response body.

Parameters:
  - method: The HTTP method for the request (e.g., "GET", "POST").
  - url: The URL to which the request is sent.
  - data: The data to include in the request body. Use nil if no data is required.
  - headers: The headers to include in the request.

Returns:
  - responseData: The response body if the request is successful. It can be a map[string]interface{} for JSON responses.
  - success: A boolean indicating whether the request was successful.
  - err: An error, if any.

“

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

“
*/
func SendRequest(method, url string, data Data, headers Header) (interface{}, bool, error) {
	// requestBody is a buffer for the request body.
	var requestBody *bytes.Buffer

	// If data is not nil, marshal it to JSON and set the request body.
	if data != nil {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return nil, false, err
		}
		requestBody = bytes.NewBuffer(jsonBytes)
	} else {
		// Otherwise, create an empty request body.
		requestBody = bytes.NewBuffer([]byte{})
	}

	// Create a new HTTP request with the given method, URL, and request body.
	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, false, err
	}

	// Add the headers to the request.
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	// Create a new HTTP client.
	client := &http.Client{}

	// Send the request and get the response.
	response, err := client.Do(request)
	if err != nil {
		return nil, false, err
	}
	defer response.Body.Close()

	// Read the response body and return it.
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, false, err
	}

	// Check if the response body is JSON.
	var dataRes map[string]interface{}
	if err := json.Unmarshal(responseBody, &dataRes); err != nil {
		// Check if the response status code is >= 400.
		if response.StatusCode >= 400 {
			return string(responseBody), false, errors.New(string(responseBody))
		} else {
			return string(responseBody), true, nil
		}
	}

	// Check if the response status code is >= 400.
	if response.StatusCode >= 400 {
		return dataRes, false, errors.New(string(responseBody))
	} else {
		return dataRes, true, nil
	}
}
