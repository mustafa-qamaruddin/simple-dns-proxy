package server

import (
	"encoding/binary"
	"github.com/mustafa-qamaruddin/simple-dns-proxy/cloudflare"
	"github.com/mustafa-qamaruddin/simple-dns-proxy/common"
	"github.com/mustafa-qamaruddin/simple-dns-proxy/custom-errors"
	dns_serde "github.com/mustafa-qamaruddin/simple-dns-proxy/dns-enc-dec"
	"github.com/pkg/errors"
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
	// Messages sent over TCP connections use server port 53 (decimal).  The
	// message is prefixed with a two byte length field which gives the message
	// length, excluding the two byte length field.  This length field allows
	// the low-level processing to assemble a complete message before beginning
	// to parse it.
	length := binary.BigEndian.Uint16(buffer[0:2])
	if length == 0 {
		custom_errors.HandleErrors(errors.New("Invalid length"), common.Error{
			Code:    101,
			Status:  custom_errors.REFUSED,
			Message: "Invalid tcp packet length",
		})
	}
	dnsQuestion, err := dns_serde.DeserializeDnsQuestion(buffer[2:])
	cloudflare.QueryDNS(dnsQuestion)
	// respond
	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	conn.Write([]byte("Hi back!"))
	conn.Write([]byte(time))
	// close conn
	conn.Close()
}
