mc - Middleware Chainer
=======================

Small package provides a simple way to chain http handlers.

[![GoDoc](https://godoc.org/github.com/aptpod/mc?status.svg)](https://godoc.org/github.com/aptpod/mc)
[![Build Status](https://travis-ci.org/aptpod/mc.svg?branch=master)](https://travis-ci.org/aptpod/mc)
[![Coverage Status](https://coveralls.io/repos/github/aptpod/mc/badge.svg?branch=master)](https://coveralls.io/github/aptpod/mc?branch=master)

Installation
------------

```
$ go get -u github.com/aptpod/mc
```

Example
-------

```
func main() {
	common := []func(http.Handler) http.Handler{
		middleware.Logger(),
		middleware.Recovery(),
	}

	authn := middleware.Authn()

	// `common` only
	http.Handle("/signin", mc.Chain(http.HandlerFunc(handler.SignIn),
		common,
	))

	// `common` and `authn`
	http.Handle("/user", mc.Chain(http.HandlerFunc(handler.User),
		common, authn,
	))

	http.ListenAndServe("127.0.0.1:8080", nil)
}
```
