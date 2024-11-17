package main

import (
	"log"
	"net"
	"os"
)

const (
	address    = "0.0.0.0:6379"
	bufferSize = 128
)

func main() {

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
		os.Exit(1)
	}
	defer l.Close()
	log.Printf("Listening on %s", address)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Accepted connection from %s", conn.RemoteAddr())

	buf := make([]byte, bufferSize)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}

	log.Printf("Read command: %s", buf[:n])

	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
	}

}
