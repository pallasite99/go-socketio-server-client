# go-socketio-server-client
Kickstarter example to develop a golang based socket.io client and server

## Dependencies
* golang needs to be installed
* currently this runs on a local golang server which a golang client connects to

### Output

1. Client side

```terminal
PS C:\Users\Dell\OneDrive\Documents\GitHub\golang-socketio-client> go run .\go-client.go
Connected to WebSocket server
Received message: Hello, Server!
```

2. Server side

```terminal
PS C:\Users\Dell\OneDrive\Documents\GitHub\golang-socketio-client> go run .\go-server.go
WebSocket server is running on ws://localhost:8080
Client connected
```

