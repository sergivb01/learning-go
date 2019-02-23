package main

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

type queryResponse struct {
	Status    string `json:"status"`
	Available bool   `json:"available"`
	Tier      string `json:"tier"`
}

var query = fmt.Sprintf("https://domain-registry.appspot.com/check?domain=%s", *domain)

func run() {
	resp, err := grequests.Get(query, nil)
	if err != nil {
		log.Println(err)
	}

	var r queryResponse
	if err := resp.JSON(&r); err != nil {
		log.Println(err)
	}

	fmt.Println(r)
}
