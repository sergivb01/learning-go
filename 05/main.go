package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	localAddr  = flag.String("lhost", ":3001", "Local address and port the proxy should listen to.")
	targetAddr = flag.String("thost", ":3000", "Remote address and port the proxy should target to.")
	logFile    = flag.String("log", "", "Set the logging file. Blank to not log into a log file.")
)

func main() {
	// parse flags and check if no args are provided
	// if so, show help and exit
	/* if flag.Parse(); len(flag.Args()) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	} */flag.Parse()

	// check if using a log file
	if *logFile != "" {
		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %s", err.Error())
		}
		defer f.Close()

		log.SetOutput(f)
		fmt.Printf("Logging to %s\n", *logFile)
	}

	// define new server configuration based on flags
	server := Server{
		Addr:   *localAddr,
		Target: *targetAddr,
	}

	// listen and serve the new server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
