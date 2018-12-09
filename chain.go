// Package mc provides a simple way to chain http handlers.
package mc

import (
	"fmt"
	"net/http"
)

// Chain returns the http.Handler chained the input middlewares.
//
// `Chain(hdl, []func(http.Handler) http.Handler{mw1, mw2}, mw3)`
// is equivalent to `mw1(mw2(mw3(hdl)))`.
//
// The function panics if the input interface is not either
// `func(http.Handler) http.Handler` or its slice.
func Chain(h http.Handler, vs ...interface{}) http.Handler {
	var ms []func(http.Handler) http.Handler

	for _, v := range vs {
		switch m := v.(type) {
		case func(http.Handler) http.Handler:
			ms = append(ms, m)
		case []func(http.Handler) http.Handler:
			ms = append(ms, m...)
		default:
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}

	return doChain(h, ms...)
}

func doChain(h http.Handler, ms ...func(http.Handler) http.Handler) http.Handler {
	if len(ms) == 0 {
		return h
	}

	return ms[0](doChain(h, ms[1:]...))
}
