package tcpserver

import (
	"fmt"
	"net"
)

// Client Represents a client with state
type Client struct {
	Quit   chan bool
	conn   net.Conn
	buffer []byte
}

// NewClient creates a new client
func NewClient(conn net.Conn) *Client {
	return &Client{
		Quit:   make(chan bool),
		conn:   conn,
		buffer: make([]byte, 1024),
	}
}

// Close the client
func (c *Client) Close() {
	fmt.Printf("[ii] Closing client '%s'\n", c.conn.RemoteAddr())
	c.Quit <- true
	c.conn.Close()
}

// Handle handles client specific logic
func (c *Client) Handle() {

	for {
		pkt, err := ReadPktBytes(c.conn)
		if err != nil {
			break
		}

		switch pkt[0] {
		case 0x30:
			c.Close()
			return
		case 0x20:
			c.conn.Write([]byte("Hello world!\n"))
		}
	}

	c.Close()
}
