package router

import (
	"bonaserver/pkg/types"
	"net/http"
	"slices"
)

type Route struct {
	Middlewares []types.Middleware
	Handler     types.NextFunc
	Path        string
}

type Router struct {
	Routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) On(p string, h types.NextFunc, m ...types.Middleware) {
	newRoute := Route{
		Middlewares: m,
		Handler:     h,
		Path:        p,
	}

	r.Routes = append(r.Routes, newRoute)
}

func SpawnMiddlewaresChain(route *Route, req *http.Request, rw http.ResponseWriter) int {
	next := func(r *http.Request, w http.ResponseWriter) {
		if len(route.Middlewares) == 1 {
			route.Handler(r, w)
			return
		}

		route.Middlewares = slices.Delete(route.Middlewares, 0, 1)
		SpawnMiddlewaresChain(route, r, w)
	}

	route.Middlewares[0](next, req, rw)
	return 0
}

func (r *Router) InitRoutes(m *http.ServeMux) {
	for _, v := range r.Routes {
		m.HandleFunc(v.Path, func(w http.ResponseWriter, r *http.Request) {
			SpawnMiddlewaresChain(&v, r, w)
		})
	}
}
