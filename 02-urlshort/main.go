package main

import (
	"fmt"
	"github.com/artslob/gophercises/02-urlshort/handlers"
	"net/http"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			_, _ = fmt.Fprintln(w, "This is url shortener. Try change path.")
		} else {
			_, _ = fmt.Fprintln(w, "Unknown path.")
		}
	})
	mux.HandleFunc("/test-result", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Seems like you were redirected from '/test', ha?")
	})
	return mux
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/test":           "/test-result",
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	//	yaml := `
	//- path: /urlshort
	// url: https://github.com/gophercises/urlshort
	//- path: /urlshort-final
	// url: https://github.com/gophercises/urlshort/tree/solution
	//`
	//	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	//	if err != nil {
	//		panic(err)
	//	}
	fmt.Println("Starting the server on http://127.1:8080")
	_ = http.ListenAndServe(":8080", mapHandler)
}
