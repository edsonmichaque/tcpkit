package tcpkit

import (
	"context"
	"io"
)

type TCPServerFunc func(req *Request) (*Response, error)

func (h TCPServerFunc) ServeTCP(r *Request) (*Response, error) {
	return h(r)
}

type TCPServer interface {
	ServeTCP(r *Request) (*Response, error)
}

type Decoder interface {
	Decode(data []byte, v interface{}) error
}

type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}

func NewServer() *Server {
	return &Server{}
}

type Server struct {
	Handler TCPServer
}

func (m *Server) ServeTCPFunc(t TCPServer) {
	m.Handler = t
}

type Request struct {
	ctx  context.Context
	body io.Reader
}

func (r Request) GetBody() io.Reader {
	return r.body
}

type Response struct {
	body io.Reader
}

func (r Response) GetBody() io.Reader {
	return r.body
}
