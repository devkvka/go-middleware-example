package main

import (
	"flag"
	"log"
	"fmt"
	"net/http"

	"mwtest/internal/handler"
	"mwtest/internal/middleware"
)

func main() {
	port := flag.String("port", ":9499", "port to serve")
	flag.Parse()

	http.HandleFunc("/", handler.Root)
	http.HandleFunc("/secure", middleware.With(
		handler.Secure,
		middleware.Auth,
	))
	http.HandleFunc("/something", middleware.With(
		handler.Something,
		middleware.Auth,
		middleware.Log,
	))

	fmt.Println("Listening on port:", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}
