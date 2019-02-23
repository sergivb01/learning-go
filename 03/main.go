package main

import (
	"net/http"
)

func main() {
	conf := Config{}

	err := LoadConfig("config.yaml", &conf)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Marshall Config struct into JSON - handle error and return json
	//json, err := json.Marshal()
	w.Write([]byte("OK"))
}
