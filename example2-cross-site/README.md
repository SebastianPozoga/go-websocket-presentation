# Chat Example

This application shows how to use the
[websocket](https://github.com/gorilla/websocket) package to implement a simple
web chat application. It is WSS version of [gorilla chat](https://github.com/gorilla/websocket/tree/master/examples/chat). Example prevent Cross-site WebSocket hijacking by check Origin (http header).
 

## Start

```
# crete certs
cd cert
go run certgen.go
cd ..
# run https and wss server
go run *.go
```

## Cross-site WebSocket hijacking
It is default safe (if you use gorilla). You can open or break the security by overload CheckOrigin:
```go:
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "https://127.0.0.1:8080" || origin == "https://127.0.0.1:8081" ||
			origin == "https://localhost:8080" || origin == "https://localhost:8081"
	},
}
```
Be carefully

## Links 
 * https://portswigger.net/web-security/websockets/cross-site-websocket-hijacking
 