package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Writer struct {
	conn *websocket.Conn
	ctx  context.Context
	id   int
	url  string
}

func NewWriter(ctx context.Context, dialer *websocket.Dialer, id int, url string) (w *Writer, err error) {
	w = &Writer{
		url: url,
		ctx: ctx,
		id:  id,
	}
	if w.conn, _, err = dialer.DialContext(w.ctx, w.url, http.Header{}); err != nil {
		return
	}
	return
}

func (w Writer) Close() (err error) {
	return w.conn.Close()
}

func (w Writer) Write(expected []string) (err error) {
	for i := 0; i < len(expected); i++ {
		if *debbug {
			fmt.Printf("(writer) Write: '%s'\n", expected[i])
		}
		if err = w.conn.WriteMessage(websocket.TextMessage, []byte(expected[i])); err != nil {
			return
		}
	}
	return
}
