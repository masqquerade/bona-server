package router

import (
	"bonaserver/pkg/types"
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

func (r *Router) InitRoutes() {
	r.Routes = []Route{
		{},
	}
}
