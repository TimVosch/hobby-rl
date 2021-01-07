package messages

// CloseConnection ...
type CloseConnection struct {
}

// MarshalBinary ...
func (m *CloseConnection) MarshalBinary() ([]byte, error) {

	return []byte{}, nil
}

// UnmarshalBinary ...
func (m *CloseConnection) UnmarshalBinary(data []byte) error {

	return nil
}
