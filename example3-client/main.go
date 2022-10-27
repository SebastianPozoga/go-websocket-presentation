package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	var (
		err  error
		conn *websocket.Conn
		resp *http.Response
	)
	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	header := http.Header{}
	if conn, resp, err = dialer.Dial("wss://localhost:8080/ws", header); err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("Response: %#v", resp)
	for i := 0; i < 10; i++ {
		if err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Message %d", i))); err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
