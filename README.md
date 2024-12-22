# gofiber-socket.io

- gofiber-socket.io is library an implementation of [Socket.IO](http://socket.io) in Golang, which is a realtime application framework.
- It using with web-framework [Go Fiber](https://gofiber.io)
- This library support socket.io-client version 3, 4 and only support websocket transport

# Contents

- [Install](#install)
- [Example](#example)

# Install

Install the package with:

```bash
go get github.com/doquangtan/gofiber-socket.io
```

Import it with:

```go
import "github.com/doquangtan/gofiber-socket.io"
```

and use `socketio` as the package name inside the code.

# Documents
## Server
### Constructor
#### socketio.New
```go
import (
	socketio "github.com/doquangtan/gofiber-socket.io"
	"github.com/gofiber/fiber/v2"
)

func socketIoRoute(app fiber.Router) {
	io := socketio.New()

	io.OnConnection(func(socket *socketio.Socket) {
		// ...
	})

	app.Use("/", io.Middleware)
	app.Route("/socket.io", io.Server)
}

func main() {
	app := fiber.New(fiber.Config{})
	app.Route("/", socketIoRoute)
	app.Listen(":3000")
}
```

### Events
#### Event: 'connection'
```go
io.OnConnection(func(socket *socketio.Socket) {
	// ...
})
```

### Methods
#### server.emit(eventName[, ...args])
```go
io.Emit("hello")
```
```go
io.Emit("hello", 1, "2", map[string]interface{}{"3": 4})
```

#### server.of(nsp)
```go
adminNamespace := io.Of("/admin")

adminNamespace.OnConnection(func(socket *socketio.Socket) {
	// ...
})
```

#### server.to(room)
```go
io.To("room-101").Emit("hello", "world")
```

#### server.fetchSockets()
```go
sockets := io.Sockets()
```

## Namespace
### Events
#### Event: 'connection'
Fired upon a connection from client.
```go
// main namespace
io.OnConnection(func(socket *socketio.Socket) {
	// ...
})

// custom namespace
io.Of("/admin").OnConnection(func(socket *socketio.Socket) {
	// ...
})
```

### Methods
#### namespace.emit(eventName[, ...args])
```go
io.Of("/admin").Emit("hello")
```
```go
io.Of("/admin").Emit("hello", 1, "2", map[string]interface{}{"3": 4})
```

#### namespace.to(room)
```go
adminNamespace := io.Of("/admin")

adminNamespace.To("room-101").Emit("hello", "world")

adminNamespace.To("room-101").To("room-102").Emit("hello", "world")
```

#### namespace.fetchSockets()
Returns the matching Socket instances:
```go
adminNamespace := io.Of("/admin")

sockets := adminNamespace.Sockets()
```

## Socket
### Events
#### Event: 'disconnect'
```go
io.OnConnection(func(socket *socketio.Socket) {
	socket.On("disconnect", func(event *socketio.EventPayload) {
		// ...
	})
})
```

### Methods
#### socket.on(eventName, callback)
Register a new handler for the given event.
```go
socket.On("news", func(event *socketio.EventPayload) {
	print(event.Data)
})
```
with several arguments
```go
socket.On("news", func(event *socketio.EventPayload) {
	if len(event.Data) > 0 && event.Data[0] != nil {
		print(event.Data[0])
	}
	if len(event.Data) > 1 && event.Data[1] != nil {
		print(event.Data[1])
	}
	if len(event.Data) > 2 && event.Data[2] != nil {
		print(event.Data[2])
	}
})
```
or with acknowledgement
```go
socket.On("news", func(event *socketio.EventPayload) {
	if event.Callback != nil {
		(*event.Callback)("hello", map[string]interface{}{
			"Test": "ok",
		})
	}
})
```

#### socket.join(room)
Adds the socket to the given room or to the list of rooms.
```go
io.Of("/test").OnConnection(func(socket *socketio.Socket) {
	socket.Join("room 237")

	io.To("room 237").Emit("a new user has joined the room")
})
```

#### socket.leave(room)
Removes the socket from the given room.
```go
io.Of("/test").OnConnection(func(socket *socketio.Socket) {
	socket.Leave("room 237");

	io.To("room 237").Emit("the user has left the room")
})
```
Rooms are left automatically upon disconnection.

#### socket.to(room)
```go
socket.On("room 237", func(event *socketio.EventPayload) {
 	// to one room
	socket.To("room 237").Emit("test", "hello")

	// to multiple rooms
	socket.To("room 237").To("room 123").Emit("test", "hello")
})
```

# Example

Please check more examples into folder in project for details. [Examples](https://github.com/doquangtan/gofiber-socket.io/tree/main/example)
