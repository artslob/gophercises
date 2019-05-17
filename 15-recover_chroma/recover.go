package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"runtime/debug"
)

type RecoverWrapper struct {
	next http.Handler
	// logStackTrace
	isDev         bool
	lst           bool
	sourcesRegexp *regexp.Regexp
}

func NewRecoverWrapper(next http.Handler, isDev bool, lst bool) *RecoverWrapper {
	if next == nil {
		panic("got empty handler")
	}
	r := regexp.MustCompile(`((/\S+)+\.go):(\d+)`)
	return &RecoverWrapper{next: next, isDev: isDev, lst: lst, sourcesRegexp: r}
}

func (wr RecoverWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		log.Println(r)
		if wr.lst {
			log.Println(string(debug.Stack()))
		}
		message := "Something went wrong"
		if wr.isDev {
			stack := string(debug.Stack())
			fmt.Println(wr.sourcesRegexp.FindAllString(stack, -1))
			message = fmt.Sprintf("error: %s\nstack:\n%s", r, stack)
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
