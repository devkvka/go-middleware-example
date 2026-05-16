// package handler implements simple http handlers
// for the routes in the project
package handler

import (
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling root for -- ", r.RemoteAddr)
	fmt.Fprintln(w, "-> Root called")
}

func Secure(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling secure for -- ", r.RemoteAddr)
	fmt.Fprintln(w, "-> Secure called")
}

func Something(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling something for -- ", r.RemoteAddr)
	fmt.Fprintln(w, "-> Something called")
}
