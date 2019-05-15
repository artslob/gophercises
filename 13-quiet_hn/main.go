package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const (
	API           = "https://hacker-news.firebaseio.com/v0/"
	NewStoriesURL = API + "newstories.json"
	ItemsURL      = API + "item/%d.json"
)

func main() {
	if err := http.ListenAndServe(":8080", root()); err != nil {
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
		ids := []int{}
		if err := json.Unmarshal(bodyBytes, &ids); err != nil {
			_, _ = fmt.Fprintf(w, "Could not parse response: %q", err)
		}
		_, _ = fmt.Fprintf(w, "ids %v", ids[:30])
		in := gen(ids[:30]...)
		storiesChannels := []<-chan StoryResponse{}
		for range ids {
			storiesChannels = append(storiesChannels, getStories(in))
		}
		for n := range merge(storiesChannels...) {
			fmt.Println(n.Id)
		}
	}
}

type StoryResponse struct {
	Story
	err error
}

func merge(storiesChannels ...<-chan StoryResponse) <-chan StoryResponse {
	out := make(chan StoryResponse)

	var wg sync.WaitGroup
	wg.Add(len(storiesChannels))

	for _, c := range storiesChannels {
		go func(c <-chan StoryResponse) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func getStories(in <-chan int) <-chan StoryResponse {
	out := make(chan StoryResponse)
	go func() {
		for n := range in {
			resp, err := http.Get(fmt.Sprintf(ItemsURL, n))
			if err != nil {
				out <- StoryResponse{err: err}
				continue
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if err != nil {
				out <- StoryResponse{err: err}
				continue
			}
			var story Story
			if err := json.Unmarshal(bodyBytes, &story); err != nil {
				out <- StoryResponse{err: err}
				continue
			}
			out <- StoryResponse{Story: story}
		}
		close(out)
	}()
	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
