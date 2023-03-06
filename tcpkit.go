package tcpkit

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

type TCPHandlerFunc func(req *Request) (*Response, error)

func (h TCPHandlerFunc) HandleTCP(r *Request) (*Response, error) {
	return h(r)
}

type TCPHandler interface {
	HandleTCP(r *Request) (*Response, error)
}

type Decoder interface {
	Decode(data []byte, v interface{}) error
}

type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

type Server struct {
	port    int
	Handler TCPHandler
}

func (m *Server) HandleTCPFunc(t TCPHandler) {
	m.Handler = t
}

func (m *Server) ListenServe() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", m.port))
	if err != nil {
		return err
	}

	for {

		con, err := l.Accept()
		if err != nil {
			continue
		}

		log.Println("new request")

		go func(con net.Conn) {
			for {
				resp, err := m.Handler.HandleTCP(&Request{
					body: con,
				})

				if err != nil {
					con.Close()
					break
				}

				io.Copy(con, resp.GetBody())
				if resp.Close {
					con.Close()
					break
				}
			}
		}(con)
	}
}

type Request struct {
	ctx  context.Context
	body io.Reader
}

func (r Request) GetBody() io.Reader {
	return r.body
}

type Response struct {
	Close bool
	Body  io.Reader
}

func (r Response) GetBody() io.Reader {
	return r.Body
}
