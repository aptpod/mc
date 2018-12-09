package mc

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func hdl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hdl"))
}

func mw(str string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(str))
			next.ServeHTTP(w, r)
		})
	}
}

func Test_Chain(t *testing.T) {
	cases := []http.Handler{
		Chain(http.HandlerFunc(hdl), mw("mw1,"), mw("mw2,"), mw("mw3,")),
		Chain(http.HandlerFunc(hdl), []func(http.Handler) http.Handler{mw("mw1,"), mw("mw2,")}, mw("mw3,")),
	}

	exp := "mw1,mw2,mw3,hdl"

	for i, c := range cases {
		h := c

		rec := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		h.ServeHTTP(rec, req)

		act := rec.Body.String()

		if act != exp {
			t.Errorf("\n   index: %d\n  actual: %#v\nexpected: %#v\n", i, act, exp)
		}
	}
}

func Test_Chain_panic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("expected panic but didn't\n")
		}
	}()

	Chain(http.HandlerFunc(hdl), mw("mw,"), 12345)
}
