package dns_packets

import (
	"encoding/binary"
	"errors"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func DecodeTcpPackets(bytes []byte) (*layers.DNS, error) {
	// Messages sent over TCP connections use server port 53 (decimal).  The
	// message is prefixed with a two byte length field which gives the message
	// length, excluding the two byte length field.  This length field allows
	// the low-level processing to assemble a complete message before beginning
	// to parse it.
	length := binary.BigEndian.Uint16(bytes[0:2])
	if length == 0 {
		return nil, errors.New("invalid tcp packet format")
	}
	return decodeDnsPackets(bytes[2:])
}

func decodeDnsPackets(bytes []byte) (*layers.DNS, error) {
	layer := CreateLayer(bytes)
	dnsLayer, ok := layer.(*layers.DNS)
	if !ok {
		return nil, errors.New("failed to decode DNS question")
	}
	return dnsLayer, nil
}

func EncodeDNSResponse(dnsLayer *layers.DNS) ([]byte, error) {
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: false,
	}
	err := dnsLayer.SerializeTo(buf, opts)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func AddError(message *layers.DNS, dnsError layers.DNSResponseCode) {
	message.ResponseCode = dnsError
}

func CreateLayer(bytes []byte) gopacket.Layer {
	var packet gopacket.Packet
	if len(bytes) == 0 {
		return nil
	}
	packet = gopacket.NewPacket(bytes, layers.LayerTypeDNS, gopacket.Default)
	return packet.Layer(layers.LayerTypeDNS)
}

func CreateErrorResponse(dnsError layers.DNSResponseCode) ([]byte, error) {
	dnsLayer := CreateLayer(make([]byte, 0)).(*layers.DNS)
	AddError(dnsLayer, dnsError)
	return EncodeDNSResponse(dnsLayer)
}
