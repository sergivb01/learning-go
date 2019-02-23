package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	apiURL string
)

// Users defines an array of User data
type Users []struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

// Address defines the address of a user
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

// Geo defines the geolocation of an address
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

// Company defines the company a user works in
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

// RunJSON executes this file code
func RunJSON() {
	apiURL = "https://jsonplaceholder.typicode.com/users"

	// Prepare the http request
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		panic(err)
	}

	// Create the http client and run the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read the body, defer closing
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var users Users
	//Unmarshal json from the response into users pointer
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}

	fmt.Printf("My JSON is %s\n", users)

	fmt.Println(len(users))
}
