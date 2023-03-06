package main

import (
	"bytes"
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

func (h enrollmentHandler) Enroll(resp *tcpkit.Response, req *tcpkit.Request) {
	buf := make([]byte, 1024)

	for i := 0; i < 25; i++ {
		_, err := req.GetBody().Read(buf)
		if err != nil {
			return
		}

		log.Println(string(buf))

		var reply string

		in := bytes.Trim(buf, "\n")
		if strings.HasPrefix(string(in), "PING") {
			reply = "PONG"
		}

		if strings.HasPrefix(string(in), "PONG") {
			reply = "PING"
		}

		_, err = resp.Write([]byte(reply + "\n"))
		if err != nil {
			return
		}
	}
}

func Logger(next tcpkit.TCPHandler) tcpkit.TCPHandler {
	return tcpkit.TCPHandlerFunc(func(resp *tcpkit.Response, req *tcpkit.Request) {
		date := time.Now()

		defer func() {
			dur := time.Now().Sub(date)
			log.Println("Duration", dur)
		}()

		next.HandleTCP(resp, req)
	})
}
