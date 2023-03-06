package tcpkit

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

type TCPHandlerFunc func(resp *Response, req *Request)

func (h TCPHandlerFunc) HandleTCP(resp *Response, req *Request) {
	h(resp, req)
}

type TCPHandler interface {
	HandleTCP(rep *Response, r *Request)
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

func (m *Server) HandleTCPFunc(t TCPHandlerFunc) {
	m.Handler = t
}

func (m *Server) HandleTCP(t TCPHandler) {
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
			defer con.Close()

			resp := Response{
				writer: con,
			}

			m.Handler.HandleTCP(&resp, &Request{
				body: con,
			})
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
	writer io.Writer
}

func (resp Response) Write(data []byte) (int, error) {
	return resp.writer.Write(data)
}
