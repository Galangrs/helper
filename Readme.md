## Installation

```
go get github.com/Galangrs/helper
```

## Example Code

### Fetching

```
package main

import (
	"fmt"

	fetch "github.com/Galangrs/helper/fetch"
)

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

	var exampleResponseInstance ExampleResponse
	response := fetch.SendRequest("POST", "http://localhost:8080/categories", data, headers, &exampleResponseInstance)
	if response.Err != nil {
		fmt.Println("success", response.StatusCode)
		fmt.Println("err", response.Err)
		return
	}

	fmt.Println("success", response.StatusCode)
}

```

### Postgres

```
package main

import (
	"log"

	"github.com/Galangrs/helper/postgres"
)

func main() {
	db, err := DBConnect("127.0.0.0","5432","postgres","postgres","postgres"")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
}
```

### Websocket

```
package main

import (
	"fmt"

	"github.com/Galangrs/helper/websocket"
)

func handleMessage(userAll websocket.UserAll, user websocket.User, msg interface{}) {
	fmt.Printf("Received message in main: %s\n", msg)
	for _, client := range userAll {
		if client != user {
			websocket.BroadcastMessage(client, msg)
		}
	}
}

func main() {
	port := "3000"
	path := "ws"

	err := websocket.ConnectWebSocket(port, path, handleMessage)
	if err != nil {
		fmt.Printf("Failed to start WebSocket server: %v\n", err)
	}

	select {}
}

```
