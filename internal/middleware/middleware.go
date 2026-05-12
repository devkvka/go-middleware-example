package middleware

import (
	"net/http"
	"slices"
)

// With applies each middleware function in the order they are passed.
// For example if you pass in Auth then Log, the authorization middleware
// will be run before any logging happens.
func With(
	h http.HandlerFunc,
	middlewares ...func(http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	for _, mw := range slices.Backward(middlewares) {
		h = mw(h)
	}
	return h
}
