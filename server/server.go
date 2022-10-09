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
	PORT = "53"
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
		logrus.Infof("start server loop %s", time.Now().Format("2 Jan 2006 15:04:05"))

		conn, err := listen.Accept()
		logrus.Infof("listen accept %s", time.Now().Format("2 Jan 2006 15:04:05"))

		if err != nil {
			custom_errors.HandleErrors(err, common.Error{
				Code:    101,
				Status:  custom_errors.REFUSED,
				Message: "Handling request",
			})
		}
		// This new goroutine will execute concurrently with the calling one.
		go handleIncomingRequest(conn)
	}
	return nil
}

func handleIncomingRequest(conn net.Conn) {
	logrus.Infof("received new request at %s", time.Now().Format("2 Jan 2006 15:04:05"))
	// store incoming data
	buffer := make([]byte, 4096)
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
	if len(packets.Questions) == 0 {
		custom_errors.HandleErrors(err, common.Error{
			Code:    101,
			Status:  custom_errors.REFUSED,
			Message: "Empty questions",
		})
	}
	reply, err := cloudflare.QueryDNS(buffer)
	if err != nil {
		custom_errors.HandleErrors(err, common.Error{
			Code:    101,
			Status:  custom_errors.REFUSED,
			Message: "Failed to read request body",
		})
	}
	// respond
	conn.Write(reply)
	// close conn
	conn.Close()
}
