package http

import (
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// A Server responds to any HTTP requests sent to NaaSgul.
type Server struct {
	Handler  http.Handler
	Listener net.Listener
	Addr     string
}

func NewServer() *Server {
	return &Server{
		Handler: httprouter.New(),
	}
}

func initRoutes() {

}
