package messages

// Ping close
type Ping struct{}

// MarshalBinary ...
func (m *Ping) MarshalBinary() ([]byte, error) {

	return []byte{}, nil
}

// UnmarshalBinary ...
func (m *Ping) UnmarshalBinary(data []byte) error {

	return nil
}
