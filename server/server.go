package server

import (
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "53"
	TYPE = "tcp"
)

func StartServer() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		// handle incoming connections
		go handleIncomingRequest(conn)
	}
}

func handleIncomingRequest(conn net.Conn) {
	// store incoming data
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	// respond
	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	conn.Write([]byte("Hi back!"))
	conn.Write([]byte(time))
	// close conn
	conn.Close()
}
