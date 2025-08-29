package main

import (
	"net"

	"github.com/bmqube/VolatileDB/handler"
	"github.com/bmqube/VolatileDB/store"
)

type Server struct {
	address string
	store   *store.Storage
	handler handler.ConnectionHandler
}

func NewServer(addr string) *Server {
	memStore := store.NewStorage()
	return &Server{
		address: addr,
		store:   memStore,
		handler: *handler.NewConnectionHandler(memStore),
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
