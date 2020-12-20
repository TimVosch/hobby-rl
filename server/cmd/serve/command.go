package serve

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/timvosch/hobby-rl/pkg/tcpserver"
	"github.com/urfave/cli"
)

// CreateCommand create the cli command for "Serve"
func CreateCommand() cli.Command {
	return cli.Command{
		Name:   "serve",
		Usage:  "Start the webserver",
		Action: serve,
	}
}

func serve(c *cli.Context) error {
	port := 3000
	addr := "0.0.0.0"
	sigChan := make(chan os.Signal, 0)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	s := tcpserver.New(addr, port)
	err := s.StartListening()
	if err != nil {
		fmt.Printf("!! Could not start server: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Server listening on %s:%d\n", addr, port)

	// Wait for exit signal
	<-sigChan
	fmt.Printf("Shutting down server...\n")
	s.Shutdown()
	return nil
}
