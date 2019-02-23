package main

import (
	"log"
	"net"
)

// Server defines server configuration
type Server struct {
	Addr   string
	Target string

	ModifyRequest  func(b *[]byte)
	ModifyResponse func(b *[]byte)
}

// ListenAndServe listens on the TCP network address laddr and then handle packets
// on incoming connections.
func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	log.Printf("Started listening on %s to target server %s", s.Addr, s.Target)

	return s.serve(listener)
}

func (s *Server) serve(ln net.Listener) error {
	for {
		// accept the new connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)

			// continue so as to skip the current connection
			continue
		}

		// handle the new connection in a new go routine
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	// connects to target server
	rconn, err := net.Dial("tcp", s.Target)
	if err != nil {
		return
	}

	// write to dst what it reads from src
	var pipe = func(src, dst net.Conn, filter func(b *[]byte)) {
		//make sure to close remote/target conn and client conn
		defer func() {
			conn.Close()
			rconn.Close()
		}()

		// make a buffer for packets
		buff := make([]byte, 65535)
		for {
			// read the connection to the buffer
			n, err := src.Read(buff)
			if err != nil {
				log.Println(err)
				return
			}

			b := buff[:n]

			if filter != nil {
				filter(&b)
			}

			_, err = dst.Write(b)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	go pipe(conn, rconn, s.ModifyRequest)
	go pipe(rconn, conn, s.ModifyResponse)
}
