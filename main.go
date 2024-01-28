package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"fmt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
	connectedClients = 0
)

func homePage(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}


func reader(conn *websocket.Conn) {
	defer func() {
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
		connectedClients--
		conn.Close()
		var connected_clients string = fmt.Sprintf("CLIENTS; %v", connectedClients)
		broadcastMessage(connected_clients)
	}()

	for {
		messageType, clientData, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// log.Println(string(clientData))
		switch string(clientData) {
		case "/connected_clients":


			var connected_clients string = fmt.Sprintf("CLIENTS; %v", connectedClients)
			sendData(messageType, connected_clients)
			log.Println("client asked for connected clients")
		default:
			log.Println(connectedClients)
			clientsMu.Lock()
			sendData(messageType, string(clientData))
			clientsMu.Unlock()
		}


	}
}

func sendData(messageType int, messageData string) {
	var data []byte
	data = append(data, messageData...)
	for client := range clients {
		if err := client.WriteMessage(messageType, data); err != nil {
			log.Println(err)
			return
		}
	}
}

func broadcastMessage(message string) {
    clientsMu.Lock()
    defer clientsMu.Unlock()

    for client := range clients {
        err := client.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            log.Println(err)
        }
    }
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	clientsMu.Lock()
	clients[ws] = true
	clientsMu.Unlock()
	connectedClients++
	log.Println("Client Successfully Connected...") 
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
