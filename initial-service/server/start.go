package server

import (
	"fmt"
	"net/http"
	"time"
)

var serverAddr = "localhost:8080"

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Feedling")
}

func Start() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayHello)

	server := &http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Started server on http://%s\n", serverAddr)
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
