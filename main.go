package main

import (
	"github.com/mustafa-qamaruddin/simple-dns-proxy/server"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	// @todo add command line args (cobra cli?)
	configs, err := NewConfigs()
	if err != nil {
		logrus.Error(errors.Wrapf(err, "Failed to read configurations"))
	}
	server.StartServer(configs)
}
