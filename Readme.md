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
