package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type RecoverWrapper struct {
	next          http.Handler
	logStackTrace bool
	isDev         bool
}

func (wr RecoverWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		log.Println(r)
		if wr.logStackTrace {
			log.Println(string(debug.Stack()))
		}
		message := "Something went wrong"
		if wr.isDev {
			message = fmt.Sprintf("error: %s\nstack:\n%s", r, debug.Stack())
		}
		http.Error(w, message, http.StatusInternalServerError)
	}()
	proxyWriter := &MiddlewareResponseWriter{ResponseWriter: w}
	wr.next.ServeHTTP(proxyWriter, r)
	_ = proxyWriter.flush()
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (w *MiddlewareResponseWriter) Write(block []byte) (int, error) {
	w.writes = append(w.writes, block)
	return len(block), nil
}

func (w *MiddlewareResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func (w *MiddlewareResponseWriter) flush() error {
	if w.status != 0 {
		w.ResponseWriter.WriteHeader(w.status)
	}
	for _, block := range w.writes {
		if _, err := w.ResponseWriter.Write(block); err != nil {
			return err
		}
	}
	return nil
}
