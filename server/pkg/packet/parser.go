package packet

import (
	"encoding"
	"encoding/binary"
	"fmt"

	"github.com/timvosch/hobby-rl/pkg/packet/messages"
)

// PktHeader contains message metadata
type PktHeader struct {
	Length uint16
}

// ByteReader ...
type ByteReader interface {
	Read(b []byte) (n int, err error)
}

func readBytes(r ByteReader, len uint) ([]byte, error) {
	buf := make([]byte, len)
	for i := 0; i < int(len); {
		cnt, err := r.Read(buf[i:])
		if err != nil {
			return nil, err
		}

		i = i + cnt
	}

	return buf, nil
}

func readPktHeader(r ByteReader) (*PktHeader, error) {
	buf, err := readBytes(r, 2)
	if err != nil {
		return nil, err
	}

	pkt := &PktHeader{}
	pkt.Length = binary.LittleEndian.Uint16(buf)
	return pkt, nil
}

// MessageFromReader ...
func MessageFromReader(r ByteReader) (interface{}, error) {
	h, err := readPktHeader(r)
	if err != nil {
		return nil, err
	}

	pktBuf, err := readBytes(r, uint(h.Length))
	if err != nil {
		return nil, err
	}

	// Instantiate correct message based on ID
	var msg encoding.BinaryUnmarshaler
	switch pktBuf[0] {
	case messages.CloseConnectionID:
		msg = &messages.CloseConnection{}
		break
	case messages.PingID:
		msg = &messages.Ping{}
		break
	case messages.HandshakeID:
		msg = &messages.Handshake{}
		break
	case messages.SendRGBID:
		msg = &messages.SendRGB{}
		break
	default:
		return nil, fmt.Errorf("No message available with id: %d", pktBuf[0])
	}

	// Unmarshal to message
	err = msg.UnmarshalBinary(pktBuf[1:])
	if err != nil {
		return nil, err
	}
	return msg, nil
}
