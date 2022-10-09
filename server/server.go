package server

import (
	"github.com/mustafa-qamaruddin/simple-dns-proxy/handler"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

const (
	HOST = "0.0.0.0"
	PORT = "53"
	TYPE = "tcp"
)

func StartServer() error {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		return err
	}
	defer listen.Close()
	logrus.Infof("Server: %s started listening for %s traffic on port %s", HOST, TYPE, PORT)
	for {
		conn, err := listen.Accept()
		logrus.Infof("Server: request received @ %s", time.Now().Format("2 Jan 2006 15:04:05"))
		if err != nil {
			logrus.Error(errors.Wrap(err, "Server: failed to accept new request"))
		}
		// This new goroutine will execute concurrently with the calling one.
		go handler.HandleIncomingRequest(conn)
	}
}
