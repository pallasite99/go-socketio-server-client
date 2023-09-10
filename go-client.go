package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	// Define the WebSocket server URL
	serverURL := "ws://localhost:8080/ws"

	// Parse the URL
	u, err := url.Parse(serverURL)
	if err != nil {
		log.Fatal(err)
	}

	// Establish a WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to WebSocket server")

	// Start a goroutine to read messages from the server
	go readMessages(conn)

	// Send a message to the server
	message := "Hello, Server!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	// Keep the program running
	select {}
}

func readMessages(conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		if messageType == websocket.TextMessage {
			handleMessage(msg)
		}
	}
}

func handleMessage(msg []byte) {
	fmt.Println("Received message:", string(msg))
}
