package middleware

import (
	"log"
	"net/http"
)

func Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logger before call -- ", r.RemoteAddr)
		h(w, r)
		log.Println("Logger after call -- ", r.RemoteAddr)
	}
}
