package main

import (
	"flag"
	"fmt"
	"github.com/artslob/gophercises/03-cyoa/story"
	"io/ioutil"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := flag.String("filename", "story.json", "json file that contains stories to parse")
	fmt.Println("Reading story from:", *filename)

	content, err := ioutil.ReadFile(*filename)
	checkError(err)

	mainStory, err := story.ParseJson(content)
	checkError(err)

	fmt.Println("Story successfully parsed")
	fmt.Println(mainStory["home"].Options)
}
