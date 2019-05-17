package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	logStackTrace := flag.Bool("lst", false, "set flag if stack trace logging to stderr is required")
	isDev := flag.Bool("dev", false, "set flag if stack trace logging to page is required")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.Handle("/panic/", RecoverWrapper{
		next:          http.HandlerFunc(panicDemo),
		logStackTrace: *logStackTrace,
		isDev:         *isDev,
	})
	mux.Handle("/panic-after/", RecoverWrapper{
		next:          http.HandlerFunc(panicAfterDemo),
		logStackTrace: *logStackTrace,
		isDev:         *isDev,
	})
	ServeFiles(mux, "/files", true)
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
