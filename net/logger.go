package net

import (
	"log"
	"net/http"
	"time"
)

var Logger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("<-- %s %s", r.Method, r.RequestURI)
		start := time.Now()

		lrw := LoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode

		log.Printf("--> End %s %s (%s)", r.Method, r.RequestURI, time.Since(start))
		log.Printf("--> %d %s", statusCode, http.StatusText(statusCode))
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func LoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}