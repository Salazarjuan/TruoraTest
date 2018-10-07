package app

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
)

type Server struct {
	port string
	Db   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing server...")
	s.port = ":" + port
	s.Db = db

}

func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)
	http.ListenAndServe(s.port, nil)
}
