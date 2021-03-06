package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	API           = "https://hacker-news.firebaseio.com/v0/"
	NewStoriesURL = API + "newstories.json"
	ItemsURL      = API + "item/%d.json"
)

var index = template.Must(template.ParseFiles("index.html"))

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
		start := time.Now()
		stories, err := getCachedStories()
		context := struct {
			Stories []StoryResponse
			Time    string
			Err     error
		}{
			Stories: stories,
			Time:    fmt.Sprintf("%.2f", time.Since(start).Seconds()),
			Err:     err,
		}
		if err := index.Execute(w, context); err != nil {
			log.Print(err)
		}
	}
}

var (
	cacheStories        = []StoryResponse{}
	cacheExpirationTime time.Time
	cacheMux            sync.Mutex
)

func getCachedStories() ([]StoryResponse, error) {
	cacheMux.Lock()
	defer cacheMux.Unlock()
	if time.Now().Before(cacheExpirationTime) {
		return cacheStories, nil
	}
	response, err := http.Get(NewStoriesURL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	ids := []int64{}
	if err := json.Unmarshal(bodyBytes, &ids); err != nil {
		return nil, err
	}
	in := gen(ids[:30]...)
	storiesChannels := []<-chan StoryResponse{}
	for range ids {
		storiesChannels = append(storiesChannels, getStories(in))
	}
	stories := []StoryResponse{}
	for story := range merge(storiesChannels...) {
		stories = append(stories, story)
	}
	sort.Sort(ByIdsDescendant(stories))
	cacheStories = stories
	cacheExpirationTime = time.Now().Add(15 * time.Second)
	return cacheStories, nil
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

func getStories(in <-chan int64) <-chan StoryResponse {
	out := make(chan StoryResponse)
	go func() {
		for n := range in {
			resp, err := http.Get(fmt.Sprintf(ItemsURL, n))
			if err != nil {
				out <- StoryResponse{Err: err}
				continue
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if err != nil {
				out <- StoryResponse{Err: err}
				continue
			}
			var story Story
			if err := json.Unmarshal(bodyBytes, &story); err != nil {
				out <- StoryResponse{Err: err}
				continue
			}
			out <- StoryResponse{Story: story}
		}
		close(out)
	}()
	return out
}

func gen(nums ...int64) <-chan int64 {
	out := make(chan int64)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
