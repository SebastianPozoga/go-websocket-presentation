package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
)

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

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", serveHome)
		mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			serveWs(hub, w, r)
		})
		server := &http.Server{
			Addr:      ":8080",
			Handler:   mux,
			TLSConfig: &tls.Config{},
		}
		if err = server.ListenAndServeTLS("./cert/cert.pem", "./cert/key.pem"); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHome)
	server := &http.Server{
		Addr:      ":8081",
		Handler:   mux,
		TLSConfig: &tls.Config{},
	}
	if err = server.ListenAndServeTLS("./cert/cert.pem", "./cert/key.pem"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
