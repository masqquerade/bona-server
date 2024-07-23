package types

import "net/http"

type NextFunc func(http.ResponseWriter, *http.Request)
type Middleware func(NextFunc, http.ResponseWriter, *http.Request)
