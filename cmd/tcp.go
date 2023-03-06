package main

import (
	"log"
	"strings"
	"time"

	"github.com/edsonmichaque/tcpkit"
)

func main() {
	h := enrollmentHandler{}

	srv := tcpkit.NewServer(12345)

	srv.HandleTCP(Logger(tcpkit.TCPHandlerFunc(h.Enroll)))
	log.Fatal(srv.ListenServe())
}

type enrollmentHandler struct{}

func (h enrollmentHandler) Enroll(req *tcpkit.Request) (*tcpkit.Response, error) {
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

func Logger(next tcpkit.TCPHandler) tcpkit.TCPHandler {
	return tcpkit.TCPHandlerFunc(func(req *tcpkit.Request) (*tcpkit.Response, error) {
		date := time.Now()

		defer func() {
			dur := time.Now().Sub(date)
			log.Println("Duration", dur)
		}()

		resp, err := next.HandleTCP(req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	})
}
