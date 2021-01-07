package messages

import "fmt"

// Handshake ...
type Handshake struct {
	UUID string
}

// MarshalBinary ...
func (m *Handshake) MarshalBinary() ([]byte, error) {

	return []byte{}, nil
}

// UnmarshalBinary ...
func (m *Handshake) UnmarshalBinary(data []byte) error {
	if len(data) != 36 {
		return fmt.Errorf("invalid message length")
	}
	m.UUID = string(data[0:36])
	return nil
}
