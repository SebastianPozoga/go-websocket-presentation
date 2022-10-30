package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Reader struct {
	conn *websocket.Conn
	ctx  context.Context
	id   int
	url  string
}

func NewReader(ctx context.Context, dialer *websocket.Dialer, id int, url string) (r *Reader, err error) {
	r = &Reader{
		url: url,
		ctx: ctx,
		id:  id,
	}
	if r.conn, _, err = dialer.DialContext(r.ctx, r.url, http.Header{}); err != nil {
		err = fmt.Errorf("(reader %d) %w", id, err)
		return
	}
	return
}

func (r Reader) Close() {
	r.conn.Close()
}

func (r Reader) Read(expected []string) (err error) {
	i := 0
	for i < len(expected) {
		var (
			messageType int
			p           []byte
			message     string
		)
		if messageType, p, err = r.conn.ReadMessage(); err != nil {
			return
		}
		if messageType != websocket.TextMessage {
			return fmt.Errorf("(reader %d) expected text message", r.id)
		}
		message = string(p)
		lines := strings.Split(message, "\n")
		for len(lines) > 0 {
			line := lines[0]
			if line != expected[i] {
				return fmt.Errorf("(reader %d) expected '%s' and take '%s'", r.id, expected[i], line)
			}
			if *debbug {
				fmt.Printf("(reader %d) line %d '%s' is ok \n", r.id, i, line)
			}
			lines = lines[1:]
			i++
		}
	}
	return
}
