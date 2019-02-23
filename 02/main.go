package main

import (
	"fmt"
	"net/http"
)

const (
	port = 80
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fmt.Println("HTTP server is starting in port 8080!")
	http.ListenAndServe(":8080", nil)
}
