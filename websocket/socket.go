package websocket

import (
	"fmt"      // Used to format the DSN string.
	"log"      // Used for logging.
	"net/http" // Used to start the WebSocket server.
	"sync"     // Used to lock the connections map.

	"golang.org/x/net/websocket" // Used for WebSocket connections.
)

/*
ConnectWebSocket establishes a WebSocket connection to a server.
Parameters:
  - port: The port of the server.
  - path: The path of the server.
  - handleMessage: The function to handle messages.

Returns:
  - err: An error, if any.
/*

“

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

“
*/

var (
	connections      = make(map[*websocket.Conn]bool)
	connectionsMutex sync.Mutex
)

type User *websocket.Conn
type UserAll []User

// WebSocketHandler handles WebSocket connections
func WebSocketHandler(ws *websocket.Conn, cb func(UserAll, User, interface{})) {
	// Lock the mutex when accessing the connections map
	connectionsMutex.Lock()
	connections[ws] = true
	connectionsMutex.Unlock()

	defer func() {
		// Lock the mutex when accessing the connections map
		connectionsMutex.Lock()
		delete(connections, ws)
		connectionsMutex.Unlock()

		// Close the WebSocket connection
		ws.Close()
	}()

	for {
		// Read message from the WebSocket connection
		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println(err)
			break
		}

		// Convert connections map to UserAll slice
		var users UserAll
		for user := range connections {
			users = append(users, user)
		}

		cb(users, ws, msg)
	}
}

func BroadcastMessage(user User, msg interface{}) {
	// Lock the mutex when accessing the connections map
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	err := websocket.Message.Send(user, msg)
	if err != nil {
		log.Println(err)
	}
}

// ConnectWebSocket starts the WebSocket server on the specified port
func ConnectWebSocket(port string, path string, callBackFunction func(UserAll, User, interface{})) error {
	// Define the WebSocket route
	http.Handle(fmt.Sprintf("/%s", path), websocket.Handler(func(ws *websocket.Conn) {
		WebSocketHandler(ws, callBackFunction)
	}))

	// Start the WebSocket server on the specified port
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		if err != nil {
			log.Fatalf("Failed to start WebSocket server on port %s: %v", port, err)
		}
	}()

	return nil
}
