package main

import (
	"log"
	"strings"

	"github.com/edsonmichaque/tcpkit"
)

func main() {
	h := handler{}

	srv := tcpkit.NewServer(12345)

	srv.HandleTCPFunc(h.Enroll)
	log.Fatal(srv.ListenServe())
}

type handler struct{}

func (h handler) Enroll(req *tcpkit.Request) (*tcpkit.Response, error) {
	buf := make([]byte, 1024)

	_, err := req.GetBody().Read(buf)
	if err != nil {
		return nil, err
	}

	log.Println(string(buf))

	return &tcpkit.Response{
		Body:  strings.NewReader("PONG\n"),
		Close: true,
	}, nil
}
