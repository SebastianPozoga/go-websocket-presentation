# Chat Example

This application shows how to use the
[websocket](https://github.com/gorilla/websocket) package to implement a simple
web chat application. It is WSS version of [gorilla chat](https://github.com/gorilla/websocket/tree/master/examples/chat)

## Start

```
# crete certs
cd cert
go run certgen.go
cd ..
# run https and wss server
go run *.go
```