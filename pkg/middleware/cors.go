package middleware

import (
	"bonaserver/pkg/types"
	"fmt"
	"net/http"
)

func SetCorsAuthHeadersMiddleware(next types.NextFunc, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://192.168.0.102:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	if r.Method == http.MethodOptions {
		fmt.Println("preflight")
		return
	}

	next(w, r)
}
