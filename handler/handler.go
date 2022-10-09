package handler

import (
	"github.com/mustafa-qamaruddin/simple-dns-proxy/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/dns/dnsmessage"
	"net"
)

func HandleIncomingRequest(conn net.Conn) {
	// close conn after done
	defer conn.Close()
	// store incoming data
	buffer := make([]byte, 4096)
	_, err := conn.Read(buffer)
	if err != nil {
		logrus.Error(errors.Wrap(err, "Handler: failed to read bytes from request"))
		HandleError(conn, dnsmessage.RCodeFormatError)
		return
	}
	reply, err := client.QueryDNS(buffer)
	if err != nil {
		logrus.Error(errors.Wrap(err, "Handler: failed to query dns info from external provider"))
		HandleError(conn, dnsmessage.RCodeServerFailure)
		return
	}
	// respond
	n, err := conn.Write(reply)
	if err != nil {
		logrus.Error(errors.Wrap(err, "Handler: failed to write dns response to connection"))
		HandleError(conn, dnsmessage.RCodeServerFailure)
		return
	}
	logrus.Infof("Handler: wrote (%d bytes)", n)
}

func createMessageWithError(errorCode dnsmessage.RCode) dnsmessage.Message {
	message := dnsmessage.Message{}
	message.Response = true
	message.RCode = errorCode
	return message
}

func writeResponse(conn net.Conn, response []byte) {
	n, err := conn.Write(response)
	if err != nil {
		logrus.Error(errors.Wrap(err, "Handler: failed to write response to connection"))
	}
	logrus.Infof("Handler: wrote (%d bytes)", n)
}

func HandleError(conn net.Conn, errorCode dnsmessage.RCode) {
	message := createMessageWithError(errorCode)
	response, err := message.Pack()
	if err != nil {
		logrus.Error(errors.Wrap(err, "Handler: failed to unpack dns message"))
	} else {
		writeResponse(conn, response)
	}
	conn.Close()
}
