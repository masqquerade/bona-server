package middleware

import (
	"bonaserver/pkg/router"
	"net/http"
	"slices"
)

func SpawnMiddlewaresChain(route *router.Route, req *http.Request, rw http.ResponseWriter) int {
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
