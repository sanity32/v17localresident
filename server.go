package v17localresident

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

const DEFAULT_RPC_SERVER_PORT int = 22217

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

type Server struct {
	addr     string
	server   *rpc.Server
	listener net.Listener
}

func (serv *Server) Setup() *Server {
	server := rpc.NewServer()
	server.RegisterName("Mouse", &MouseCtl{})
	server.RegisterName("Key", &KeyCtl{})
	server.RegisterName("Screenshot", &ScreenshotCtl{})
	server.HandleHTTP("/", "/debug")

	serv.server = server
	return serv
}

func (serv *Server) Run() (<-chan any, error) {
	if serv.server == nil {
		return nil, errors.New("rpc server is not set up")
	}
	l, err := net.Listen("tcp", serv.addr)
	if err != nil {
		return nil, err
	}

	serv.listener = l
	log.Println("Serving address", serv.addr)
	var ch chan any
	go func() {
		serv.server.Accept(l)
		ch <- nil
	}()
	return ch, nil
}
