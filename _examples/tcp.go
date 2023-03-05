package examples

import (
	"github.com/edsonmichaque/tcpkit"
	"github.com/go-playground/locales/haw"
)

func main() {
	h := handler{}

	srv := tcpkit.NewServer()

	srv.ServeTCPFunc(h.Enroll())
}

type handler struct{}

func (h handler) Enroll(req *tcpkit.Request) (*tcpkit.Response, error) {
	return nil, errors.New("not implemented")
}
