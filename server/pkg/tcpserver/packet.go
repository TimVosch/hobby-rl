package tcpserver

import (
	"encoding/binary"
	"net"
)

// PktHeader contains message metadata
type PktHeader struct {
	Length uint16
}

func readBytes(c net.Conn, len uint) ([]byte, error) {
	buf := make([]byte, len)
	for i := 0; i < int(len); {
		cnt, err := c.Read(buf[i:])
		if err != nil {
			return nil, err
		}

		i = i + cnt
	}

	return buf, nil
}

// ReadPktHeader ...
func ReadPktHeader(c net.Conn) (*PktHeader, error) {
	buf, err := readBytes(c, 2)
	if err != nil {
		return nil, err
	}

	pkt := &PktHeader{}
	pkt.Length = binary.LittleEndian.Uint16(buf)
	return pkt, nil
}

// ReadPktBytes ...
func ReadPktBytes(c net.Conn) ([]byte, error) {
	h, err := ReadPktHeader(c)
	if err != nil {
		return nil, err
	}

	pktBuf, err := readBytes(c, uint(h.Length))
	if err != nil {
		return nil, err
	}
	return pktBuf, nil
}
