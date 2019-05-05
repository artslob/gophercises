package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const API = "https://hacker-news.firebaseio.com/v0/"

var NewStoriesURL = API + "newstories.json"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", root())
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("got error: %v\n", err)
	}
}

func root() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		response, err := http.Get(NewStoriesURL)
		if err != nil {
			_, _ = fmt.Fprintf(w, "Some error occurred while processing request: %q", err)
			return
		}
		defer func() { _ = response.Body.Close() }()
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			_, _ = fmt.Fprintf(w, "Some error occurred while processing request: %q", err)
			return
		}
		body := string(bodyBytes)
		_, _ = fmt.Fprint(w, body)
	}
}
