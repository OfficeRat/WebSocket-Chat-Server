# WebSocket Chat Server

This is a simple WebSocket chat server implemented in Go, using the Gorilla WebSocket library. The server allows clients to connect via WebSocket and exchange messages in real-time. The front end is a basic HTML file (index.html) that serves as the user interface for the chat.
### Prerequisites

Go installed on your machine
Gorilla WebSocket library (github.com/gorilla/websocket) installed:

```bash
go get -u github.com/gorilla/websocket
```
### Getting Started

Clone the repository:

```bash
git clone https://github.com/OfficeRat/WebSocket-Chat-Server.git
```
Navigate to the project directory:

```bash
cd WebSocket-Chat-Server
```
Run the server:

```bash
go run main.go
```
Open a web browser and navigate to http://localhost:8080 to access the chat interface.

### Project Structure
main.go: Main entry point for the application. It sets up the HTTP server, WebSocket endpoint, and handles client connections.
 
index.html: Simple HTML file serving as the frontend for the chat.

### WebSocket Endpoint

The WebSocket endpoint is located at /ws. Clients can connect to this endpoint to establish a WebSocket connection.
WebSocket Communication

Client Connection: When a client successfully connects, the server logs the event and increments the connected clients count.

Message Handling: The server handles different types of messages received from clients. By default, it broadcasts messages to all connected clients. If a client sends the message "/connected_clients," the server responds with the current number of connected clients.

Disconnection: When a client disconnects, the server removes the client from the list of connected clients and updates the count. It also sends a broadcast message to inform other clients about the change.

### Customization

Feel free to customize the code to add features, improve security, or enhance the user interface. The server is a starting point for building a more robust WebSocket chat application.