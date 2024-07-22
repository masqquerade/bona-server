package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type authContentKey string

func AuthUserMiddleware(next http.Handler, token string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authParts := strings.Split(r.Header.Get("Authorization"), " ")
		if len(authParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			fmt.Println("len in smaller then 2")

			return
		}

		authType := authParts[0]
		authData := authParts[1]

		switch authType {
		case "tma":
			if err := initdata.Validate(authData, token, 24*time.Hour); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				fmt.Println(err.Error())

				return
			}

			initData, err := initdata.Parse(authData)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				fmt.Println(err.Error())

				return
			}

			authCtx := context.WithValue(r.Context(), authContentKey("initData"), initData)
			authR := r.WithContext(authCtx)

			next.ServeHTTP(w, authR)
		}
	})
}
