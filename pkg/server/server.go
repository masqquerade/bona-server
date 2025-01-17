package server

import (
	"bonaserver/pkg/middleware"
	"bonaserver/pkg/router"
	"bonaserver/pkg/store"
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	Token  string
	Store  *store.Store
	Router *router.Router
}

func NewServer(token string, store *store.Store) (*Server, error) {
	return &Server{
		Token: token,
		Store: store,
	}, nil
}

func (s *Server) InitServer() *http.ServeMux {
	s.mux = http.NewServeMux()
	s.Router = router.NewRouter()

	s.Router.On("/api/set-img", s.SetImgHandler, middleware.SetCorsAuthHeadersMiddleware, middleware.AuthUserMiddleware)

	s.Router.InitRoutes(s.mux)

	return s.mux
}

func (s *Server) SetImgHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server side"))
}
