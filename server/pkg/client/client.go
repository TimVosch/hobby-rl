package client

import (
	"fmt"
	"net"

	"github.com/timvosch/hobby-rl/pkg/packet"
	"github.com/timvosch/hobby-rl/pkg/packet/messages"
)

// Client Represents a client with state
type Client struct {
	ID     string
	Quit   chan bool
	conn   net.Conn
	buffer []byte
}

// New creates a new client
func New(conn net.Conn) *Client {
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

	handlerFunc := func() error {
		// Infinite loop, returns an error/nil
		for {
			msgI, err := packet.MessageFromReader(c.conn)
			if err != nil {
				return err
			}

			//
			switch msg := msgI.(type) {
			case *messages.CloseConnection:
				return nil
			case *messages.Ping:
				c.conn.Write([]byte("Hello world!\n"))
				break
			case *messages.Handshake:
				if c.ID != "" {
					return fmt.Errorf("client is setting ID while it already had an ID")
				}
				c.ID = msg.UUID
				fmt.Printf("[ii] Client identified as: %s\n", c.ID)
				break
			case *messages.SendRGB:
				break
			}
		}
	}

	// Execute the handler function and catch a possible error
	err := handlerFunc()
	if err != nil {
		fmt.Printf("[EE] Error occured: %s\n", err)
	}
	c.Close()
}
