package cloudflare

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
)

const (
	HOST            = "1.1.1.1"
	HostAlternative = "1.0.0.1"
	PORT            = "853"
	TlsHost         = "cloudflare-dns.com"
	TYPE            = "tcp"
)

type CloudFlareClient struct {
	connection tls.Conn
}

func NewCloudFlareClient() *CloudFlareClient {
	return nil
}

func QueryDNS(bytes []byte) ([]byte, error) {
	config := tls.Config{}
	conn, err := tls.Dial(TYPE, fmt.Sprintf("%s:%s", HOST, PORT), &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
		return nil, err
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		fmt.Println(v.Subject)
	}
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

	n, err := io.WriteString(conn, string(bytes))
	if err != nil {
		log.Fatalf("client: write: %s", err)
		return nil, err
	}
	log.Printf("client: wrote %q (%d bytes)", string(bytes), n)

	reply := make([]byte, 4096)
	n, err = conn.Read(reply)
	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	log.Print("client: exiting")
	return reply, nil
}
