package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var debbug = flag.Bool("debbug", false, "is debbug mode")
var readerCounter = flag.Int("readers", 60, "reader count")
var timeout = flag.Int("timeout", 60*5, "time to stop")
var url = flag.String("url", "wss://localhost:8080/ws", "websocket url")

func main() {
	flag.Parse()
	mainGo()
}

func mainGo() {
	var (
		err    error
		ctx    context.Context
		cancel context.CancelFunc

		writer *Writer

		wg = &sync.WaitGroup{}

		errsMU sync.Mutex
		errs   []error
	)

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)

	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 3 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	wg.Add(*readerCounter + 1)
	for i := 0; i < *readerCounter; i++ {
		var r *Reader
		if r, err = NewReader(ctx, dialer, i, *url); err != nil {
			panic(err)
		}
		go func(i int, r *Reader) {
			defer wg.Done()
			defer r.Close()
			var (
				rErr error
			)
			if rErr = r.Read(testData); rErr != nil {
				errsMU.Lock()
				defer errsMU.Unlock()
				errs = append(errs, rErr)
				cancel()
			}
		}(i, r)
	}
	fmt.Printf("Readers connected\n")

	if writer, err = NewWriter(ctx, dialer, 0, *url); err != nil {
		panic(err)
	}
	go func() {
		defer wg.Done()
		defer writer.Close()
		if err = writer.Write(testData); err != nil {
			panic(err)
		}
	}()
	fmt.Printf("Writer connected\n")

	wg.Wait()
	for i := 0; i < len(errs); i++ {
		fmt.Printf(" * Error (%d): %v\n", i, errs[i])
	}

	fmt.Printf("Finished\n")
	return
}
