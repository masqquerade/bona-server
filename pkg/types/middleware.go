package types

import "net/http"

type NextFunc func(*http.Request, http.ResponseWriter)
type Middleware func(NextFunc, *http.Request, http.ResponseWriter)
