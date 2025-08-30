package main

import (
	"net"

	"github.com/bmqube/VolatileDB/handlers"
	"github.com/bmqube/VolatileDB/store"
)

type Server struct {
	address string
	store   *store.Storage
	handler handlers.ConnectionHandler
}

func NewServer(addr string) *Server {
	memStore := store.NewStorage()
	return &Server{
		address: addr,
		store:   memStore,
		handler: *handlers.NewConnectionHandler(memStore),
	}
}

func (server *Server) Start() error {
	listener, err := net.Listen("tcp", server.address)

	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go server.handler.Handle(conn)
	}
}

func main() {
	server := NewServer(":6969")
	if err := server.Start(); err != nil {
		panic(err)
	}
}
