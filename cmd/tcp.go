package main

import (
	"log"
	"strings"
	"time"

	"github.com/edsonmichaque/tcpkit"
)

func main() {
	h := handler{}

	srv := tcpkit.NewServer(12345)

	srv.HandleTCPFunc(Logger(tcpkit.TCPServerFunc(h.Enroll)))
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
		Body: strings.NewReader("PONG\n"),
	}, nil
}

func Logger(next tcpkit.TCPServer) tcpkit.TCPServer {
	return tcpkit.TCPServerFunc(func(req *tcpkit.Request) (*tcpkit.Response, error) {
		date := time.Now()

		defer func() {
			dur := time.Now().Sub(date)
			log.Println("Duration", dur)
		}()

		resp, err := next.ServeTCP(req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	})
}
