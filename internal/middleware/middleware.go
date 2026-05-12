package middleware

import (
	"net/http"
	"slices"
)

func With(
	h http.HandlerFunc,
	middlewares ...func(http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	for _, mw := range slices.Backward(middlewares) {
		h = mw(h)
	}
	return h
}
