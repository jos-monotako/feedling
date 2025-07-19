package server

import (
	"fmt"
	"net/http"
)

var serverAddr = "localhost:8080"

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Feedling")
}

func Start() error {
	fmt.Printf("Started server on http://%s\n", serverAddr)
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		return err
	}

	return nil
}
