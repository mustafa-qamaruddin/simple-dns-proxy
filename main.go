package main

import (
	"github.com/mustafa-qamaruddin/simple-dns-proxy/handler"
	"github.com/mustafa-qamaruddin/simple-dns-proxy/server"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	err := server.StartServer(handler.HandleIncomingRequest)
	if err != nil {
		logrus.Error(errors.Wrapf(err, "Failed to start proxy server"))
	}
}
