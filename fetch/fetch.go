package fetch

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

/*
SendRequest sends an HTTP request and returns the response body.

Parameters:
  - method: The HTTP method for the request (e.g., "GET", "POST").
  - url: The URL to which the request is sent.
  - data: The data to include in the request body. Use nil if no data is required.
  - headers: The headers to include in the request.
  - Expect: The Response struct to unmarshal the response body into. Use nil if no response body is expected.

Returns:
  - response: {
		Status:     bool,
	    StatusCode: int,
        Body:       io.ReadCloser,
        Err:        error,
  	}

“

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

“
*/

// Header struct represents the headers for the request
type Header map[string]string

// Data struct represents the data for the request
type Data map[string]interface{}

// CallbackFunc is a callback function signature
type CallbackFunc func(interface{})

type Response struct {
	Status     bool
	StatusCode int
	Body       interface{}
	Header     http.Header
	Err        error
}

// SendRequest sends an HTTP request and returns the response body.

func SendRequest(method, url string, data Data, headers Header, expectResponse interface{}) Response {
	// requestBody is a buffer for the request body.
	var requestBody *bytes.Buffer

	// If data is not nil, marshal it to JSON and set the request body.
	if data != nil {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return Response{
				Status:     false,
				Header:     nil,
				StatusCode: 0,
				Body:       nil,
				Err:        err,
			}
		}
		requestBody = bytes.NewBuffer(jsonBytes)
	} else {
		// Otherwise, create an empty request body.
		requestBody = bytes.NewBuffer([]byte{})
	}

	// Create a new HTTP request with the given method, URL, and request body.
	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return Response{
			Status:     false,
			Header:     nil,
			StatusCode: 0,
			Body:       nil,
			Err:        err,
		}
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
		return Response{
			Status:     false,
			Header:     nil,
			StatusCode: 0,
			Body:       nil,
			Err:        err,
		}
	}
	defer response.Body.Close()

	// Read the response body.
	responseBodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{
			Status:     false,
			Header:     response.Header,
			StatusCode: int(response.StatusCode),
			Body:       nil,
			Err:        err,
		}
	}

	// If expectResponse is provided, unmarshal the response body into it.
	if expectResponse != nil {
		if err := json.Unmarshal(responseBodyByte, expectResponse); err != nil {
			return Response{
				Status:     false,
				Header:     response.Header,
				StatusCode: int(response.StatusCode),
				Body:       nil,
				Err:        err,
			}
		}
	}

	return Response{
		Status:     response.StatusCode < 400,
		Header:     response.Header,
		StatusCode: int(response.StatusCode),
		Body:       expectResponse,
		Err:        errors.New(string(responseBodyByte)),
	}
}
