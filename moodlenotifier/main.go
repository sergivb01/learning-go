package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	httpClient = &http.Client{Timeout: time.Second * 10}
	request    *http.Request
	config     struct {
		// TODO: Can't multiply duration * timeout (int/float32)
		Timeout int               `yaml:"timeout,omitempty"`
		BaseURL string            `yaml:"baseURL,omitempty"`
		Cookies map[string]string `yaml:"cookies,omitempty"`
	}
)

func main() {
	t := time.Now()

	c := make(chan bool)
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, os.Interrupt)

	// read configuration file
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error while opening config file: %v\n", err)
	}
	if err := yaml.Unmarshal(b, &config); err != nil {
		log.Fatalf("error while unmarshaling config: %v\n", err)
	}

	// create request and save it into a package var. It will always be the same
	if err = createRequest(); err != nil {
		log.Fatalf("error creating request: %v\n", err)
	}

	go checkForMark(t, signs, c)

	if ok := <-c; ok {
		log.Printf("WE HAVE A MARK!!! It took: %s\n", time.Since(t))
	}
	log.Println("it looks like I did an oopsie doopsie...")
}

// checkForMark runs the request, checks for the response and listens for chans
func checkForMark(t time.Time, s <-chan os.Signal, c chan bool) {
	for {
		select {
		case <-s:
			log.Println("recived signals... stopping")
			c <- false
			return

		default:
			res, err := runRequest() // execute the request saved in the pkg var
			if err != nil {
				log.Fatal(err)
			}

			if !strings.Contains(res, "Sense qualificació") {
				c <- true
				return
			} else if !strings.Contains(res, "Estat de la tramesa") {
				c <- false
				// panic as we have no way to continue the application and I don't know
				// how to handle it either
				panic("cookie is not (or no longer) valid")
			}

			time.Sleep(time.Second * 15)
		}
		log.Printf("we have no mark yet... Time elapsed: %s\n", time.Since(t))
	}
}

// createRequests bakes the HTTP GET request that will be executed upon the
// status fetch. It will return an error if it fails
func createRequest() error {
	req, err := http.NewRequest("GET", config.BaseURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cache-Control", "no-cache")

	for key, val := range config.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  key,
			Value: val,
		})
	}
	log.Printf("added cookies: %+v\n", req.Cookies())
	request = req

	return nil
}

// runRequest does the request created in the previous process. It will return
// an string of the body and an error
func runRequest() (string, error) {
	resp, err := httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}
	defer resp.Body.Close() // always remember to close the body.

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body: %v", err)
	}

	return string(body), nil
}
