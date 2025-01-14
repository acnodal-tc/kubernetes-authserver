package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		if r.RequestURI != "/healthz" && OPT_DEBUG {
			log.Printf(
				"%s\t%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				r.RemoteAddr,
				name,
				time.Since(start),
			)
		}

	})
}
