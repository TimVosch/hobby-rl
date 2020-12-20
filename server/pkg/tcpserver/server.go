package tcpserver

import (
	"context"
	"fmt"
	"net"
)

// Server contains configuration and state
type Server struct {
	port      int
	addr      string
	ctx       context.Context
	ctxCancel context.CancelFunc
}

// New creates a new instance of Server
func New(addr string, port int) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		addr:      addr,
		port:      port,
		ctx:       ctx,
		ctxCancel: cancel,
	}
}

// StartListening ...
func (s *Server) StartListening() error {
	l, err := net.Listen(
		"tcp4",
		fmt.Sprintf("%s:%d", s.addr, s.port),
	)
	if err != nil {
		return err
	}

	go func() {
		ctx := s.ctx
		for {
			select {
			case <-ctx.Done():
				return
			default:
				c, err := l.Accept()
				if err != nil {
					fmt.Printf("Error accepting client! %s\n", err)
					continue
				}
				fmt.Printf("Accepted client!\n")
				go ClientHandler(c)
			}
		}
	}()
	return nil
}

// Shutdown stops serving clients
func (s *Server) Shutdown() error {
	s.ctxCancel()
	return nil
}
