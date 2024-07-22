package server

import (
	"bonaserver/pkg/store"
	"net/http"
)

type Server struct {
	mux   *http.ServeMux
	Token string
	Store *store.Store
}

func NewServer(token string, store *store.Store) (*Server, error) {
	return &Server{
		Token: token,
		Store: store,
	}, nil
}

func (s *Server) InitServer() *http.ServeMux {
	s.mux = http.NewServeMux()
	s.InitRoutes()

	return s.mux
}

func (s *Server) InitRoutes() {

}