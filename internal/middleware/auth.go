package middleware

import (
	"fmt"
	"net/http"
)

// Auth is a simple authorization middleware that checks for
// an auth token in the request header
func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != "secure" {
			http.Error(w, "-> unauthorized", http.StatusUnauthorized)
			fmt.Println("Unauthed call from -- ", r.RemoteAddr)
			return
		}

		fmt.Println("authed user -- ", r.RemoteAddr)

		h(w, r)
	}
}
