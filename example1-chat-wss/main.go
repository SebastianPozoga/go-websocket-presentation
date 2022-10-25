package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http/https service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	var (
		err error
	)

	flag.Parse()
	hub := newHub()
	go hub.run()

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHome)
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	server := &http.Server{
		Addr:      *addr,
		Handler:   mux,
		TLSConfig: &tls.Config{},
	}
	// err := http.ListenAndServe(*addr, nil)
	fmt.Println("start")
	err = server.ListenAndServeTLS("./cert/cert.pem", "./cert/key.pem")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
