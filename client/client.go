package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

const (
	HOST = "1.1.1.1"
	PORT = "853"
	TYPE = "tcp"
)

func QueryDNS(bytes []byte) ([]byte, error) {
	config := tls.Config{}
	conn, err := tls.Dial(TYPE, fmt.Sprintf("%s:%s", HOST, PORT), &config)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	logrus.Infof("Client: connected to: %s", conn.RemoteAddr())

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		cert, _ := x509.MarshalPKIXPublicKey(v.PublicKey)
		logrus.Infof("Certificate: %q", cert)
		logrus.Infof("Subject: %s", v.Subject)
	}
	logrus.Infof("Client: handshake: %t", state.HandshakeComplete)

	n, err := io.WriteString(conn, string(bytes))
	if err != nil {
		return nil, err
	}
	logrus.Infof("Client: sent to external provider (%d bytes)", n)

	reply := make([]byte, 4096)
	n, err = conn.Read(reply)
	logrus.Infof("Client: read (%d bytes)", n)
	return reply, err
}
