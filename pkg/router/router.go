package router

import (
	"bonaserver/pkg/types"
	"net/http"
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

func HandleRoute(route *Route, rw http.ResponseWriter, req *http.Request, i int) {
	next := func(w http.ResponseWriter, r *http.Request) {
		if i == (len(route.Middlewares) - 1) {
			route.Handler(w, r)
		} else {
			HandleRoute(route, w, r, i+1)
		}

	}

	route.Middlewares[i](next, rw, req)
}

func (router *Router) InitRoutes(m *http.ServeMux) {
	for _, v := range router.Routes {
		m.HandleFunc(v.Path, func(w http.ResponseWriter, r *http.Request) {
			HandleRoute(&v, w, r, 0)
		})
	}
}
