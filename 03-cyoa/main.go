package main

import (
	"flag"
	"fmt"
	"github.com/artslob/gophercises/03-cyoa/story"
	"io/ioutil"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := flag.String("filename", "story.json", "json file that contains stories to parse")
	port := flag.Int("port", 8888, "the port to start the server on")
	flag.Parse()
	fmt.Println("Reading story from:", *filename)

	content, err := ioutil.ReadFile(*filename)
	checkError(err)

	mainStory, err := story.ParseJson(content)
	checkError(err)

	fmt.Println("Story successfully parsed")

	fmt.Printf("Starting the server on http://127.1:%d\n", *port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), storyHandler(mainStory))
	checkError(err)
}

func storyHandler(mainStory story.Story) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if a, ok := mainStory[path[1:]]; ok {
			_, _ = writer.Write([]byte(a.Title))
		} else {
			_, _ = writer.Write([]byte("How did you get here?"))
		}
	}
}
