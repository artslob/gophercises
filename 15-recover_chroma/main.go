package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

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

func wrapper(next http.HandlerFunc, logStackTrace bool, isDev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}
			log.Println(r)
			if logStackTrace {
				log.Println(string(debug.Stack()))
			}
			message := "Something went wrong"
			if isDev {
				message = fmt.Sprintf("error: %s\nstack:\n%s", r, debug.Stack())
			}
			http.Error(w, message, http.StatusInternalServerError)
		}()
		writer := &MiddlewareResponseWriter{ResponseWriter: w}
		next(writer, r)
		_ = writer.flush()
	}
}

func main() {
	logStackTrace := flag.Bool("lst", false, "set flag if stack trace logging to stderr is required")
	isDev := flag.Bool("dev", false, "set flag if stack trace logging to page is required")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/panic/", wrapper(panicDemo, *logStackTrace, *isDev))
	mux.HandleFunc("/panic-after/", wrapper(panicAfterDemo, *logStackTrace, *isDev))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "This is root handler")
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "panic after demo\n")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}
