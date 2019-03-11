package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	c := NewClient("https://kctbh9vrtdwd.statuspage.io/api/v2", nil)
	bytes, err := c.doRequest("GET", "/summary.json")
	if err != nil {
		log.Fatalf("could not make a complete request: %v\n", err)
		os.Exit(1)
	}

	var res StatusResponse
	if err := json.Unmarshal(bytes, &res); err != nil {
		log.Fatalf("could not marshal response: %v\n", err)
	}
	fmt.Println(res)

	// c.baseURL = "https://httpbin.org"
	// bytes, err = c.doRequest("GET", "/get")
	// if err != nil {
	// 	log.Fatalf("could not complete request: %v\n", err)
	// }
	// fmt.Println(string(bytes))

}
