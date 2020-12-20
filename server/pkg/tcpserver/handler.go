package tcpserver

import (
	"fmt"
	"net"
)

// ClientHandler handles client specific logic
func ClientHandler(c net.Conn) {
	fmt.Printf("Hello from handler\n")
	c.Close()
}
