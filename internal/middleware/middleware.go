package middleware

import (
	"net/http"
	"slices"
)

// With applies each middleware function in the order they are passed
func With(
	h http.HandlerFunc,
	middlewares ...func(http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	// With slices.Backward the middleware gets applied in the order they are
	// passed in to the function.
	for _, mw := range slices.Backward(middlewares) {
		h = mw(h)
	}
	return h
}
