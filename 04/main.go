// Simple TCP chat system using Go concurrency
// Use telnet to connect and chat @ 127.0.0.1:3000
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// create net listener using TCP protocol and listen to port 3000
	l, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer l.Close() // always defer closing the listener

	// need maps & chans for active connections, new connections,
	// dead connections and messages
	aconns := make(map[net.Conn]int)
	conns := make(chan net.Conn)
	dconns := make(chan net.Conn)
	msgs := make(chan string)

	i := 0 // used by client IDs

	// go routine to accept all connections
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println(err.Error())
			}

			log.Println(conn)
			conns <- conn
		}
	}()

	for {

		select {
		// set the connection ID
		case conn := <-conns:
			aconns[conn] = i
			i++

			go func(conn net.Conn, i int) {
				rd := bufio.NewReader(conn)

				for {
					m, err := rd.ReadString('\n')
					if err != nil {
						log.Println(err.Error())
					}

					log.Println(m)
					msgs <- fmt.Sprintf("[Client %v] %v", i, m)
				}
			}(conn, i)

		case msg := <-msgs:
			// send message to all active connections
			for conn := range aconns {
				conn.Write([]byte(msg))
			}
			log.Println(msg)

		case dconn := <-dconns:
			// Make sure to delete before sending message
			delete(aconns, dconn)

			m := fmt.Sprintf("[Client %v] - Disconnected", aconns[dconn])

			for conn := range aconns {
				conn.Write([]byte(m))
			}
			log.Println(m)
		}
	}

}
