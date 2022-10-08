package dns_enc_dec

import (
	"errors"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func DecodeDnsQuestion(bytes []byte) (*layers.DNS, error) {
	packet := gopacket.NewPacket(bytes, layers.LayerTypeDNS, gopacket.Default)
	dnsPacket := packet.Layer(layers.LayerTypeDNS)
	tcp, ok := dnsPacket.(*layers.DNS)
	if !ok {
		return nil, errors.New("failed to decode DNS question")
	}
	return tcp, nil
}
