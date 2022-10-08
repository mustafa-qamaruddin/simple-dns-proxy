package server

import (
	"github.com/mustafa-qamaruddin/simple-dns-proxy/cloudflare"
	"github.com/mustafa-qamaruddin/simple-dns-proxy/common"
	"github.com/mustafa-qamaruddin/simple-dns-proxy/custom-errors"
	"github.com/mustafa-qamaruddin/simple-dns-proxy/dns-packets"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

const (
	HOST = "localhost"
	PORT = "9953"
	TYPE = "tcp"
)

func StartServer(*common.Configs) error {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		return err
	}
	defer listen.Close()
	logrus.Infof("Server %s started listening for %s traffic on port %s", HOST, TYPE, PORT)
	for {
		conn, err := listen.Accept()
		if err != nil {
			custom_errors.HandleErrors(err, common.Error{
				Code:    101,
				Status:  custom_errors.REFUSED,
				Message: "Handling request",
			})
		}
		go handleIncomingRequest(conn)
	}
	return nil
}

func handleIncomingRequest(conn net.Conn) {
	// store incoming data
	buffer := make([]byte, 514)
	_, err := conn.Read(buffer)
	if err != nil {
		custom_errors.HandleErrors(err, common.Error{
			Code:    101,
			Status:  custom_errors.REFUSED,
			Message: "Failed to read request body",
		})
	}
	packets, err := dns_packets.DecodeTcpPackets(buffer)
	if err != nil {
		custom_errors.HandleErrors(err, common.Error{
			Code:    101,
			Status:  custom_errors.REFUSED,
			Message: "Failed to read request body",
		})
	}
	cloudflare.QueryDNS(packets)
	// respond
	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	conn.Write([]byte("Hi back!"))
	conn.Write([]byte(time))
	// close conn
	conn.Close()
}
