package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

var (
	domain = flag.String("domain", "sergi.dev", "Domain that is being checked")
)

func main() {
	if flag.Parse(); *domain == "" {
		fmt.Println("Please specify a domain")
	}

	// create a channel thought which we are going to recive the statuses
	c := make(chan bool)
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	// goroutine to run the checks
	go func() {
		for {
			checkDomain(c)
			time.Sleep(time.Second)
		}
	}()

LOOP:
	for {
		select {
		case ok := <-c:
			if !ok {
				errorf("%s is already taken! OH NOES! D;", *domain)
			}

		case <-signals:
			fmt.Fprint(os.Stderr, "quitting")
			close(c)
			break LOOP
		}
	}

}

//checkDomain checks whether a domain is available or not - ugly but working way
// returns bool if the domain is available
func checkDomain(c chan bool) {
	_, err := net.LookupIP(*domain)
	c <- err != nil
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(2)
}
