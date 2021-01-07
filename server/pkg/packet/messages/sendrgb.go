package messages

// SendRGB ...
type SendRGB struct {
	Target string
	R      int
	G      int
	B      int
}

// MarshalBinary ...
func (m *SendRGB) MarshalBinary() ([]byte, error) {

	return []byte{}, nil
}

// UnmarshalBinary ...
func (m *SendRGB) UnmarshalBinary(data []byte) error {

	return nil
}
