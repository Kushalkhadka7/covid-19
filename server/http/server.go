package http

import (
	util "covid_server/util"
	"fmt"
	"log"
	"net"
)

// Server initializes new http server and grpc server.
type Server struct {
	port int
}

// Serverer creates server.
type Serverer interface {
	CreateHTTPServer() (net.Listener, error)
	CreateGRPCServer()
	Stop()
}

// New intializes port to the server.
func New(port int) *Server {
	return &Server{
		port,
	}
}

// CreateHTTPServer creates new http server and run it on the given port.
func (s *Server) CreateHTTPServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return nil, util.Error("Unable to establish http server: ", err)
	}

	log.Printf("Server listening on port:%d", s.port)
	return listener, nil
}

// Stop server.
func (s *Server) Stop() {
	s.Stop()
}
