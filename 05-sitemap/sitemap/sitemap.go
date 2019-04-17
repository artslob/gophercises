package sitemap

import (
	"github.com/artslob/gophercises/04-link/parser"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type siteMap struct {
	locations map[string]bool
	lock      *sync.Mutex
}

func NewSiteMap() *siteMap {
	return &siteMap{
		locations: map[string]bool{},
		lock:      &sync.Mutex{},
	}
}

func (s *siteMap) GetMap() map[string]bool {
	return s.locations
}

// BuildMap requires origin to be url without schema and trailing slash.
// Supports only links starting with '//' or '/' or 'https://'.
func (s *siteMap) BuildMap(origin string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go s.buildMap(origin, "https://"+origin, &wg)
	wg.Wait()
}

func (s *siteMap) buildMap(origin, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	if strings.HasPrefix(url, "//") {
		url = "https:" + url
	}
	if strings.HasPrefix(url, "/") {
		url = "https://" + origin + url
	}
	if !strings.HasPrefix(url, "https://"+origin) {
		return
	}
	if !s.add(url) { // means url already taken into account
		return
	}
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	_ = res.Body.Close()
	walker := parser.NewWalker(string(bytes))
	for _, link := range walker.Links {
		wg.Add(1)
		go s.buildMap(origin, link.Href, wg)
	}
}

func (s *siteMap) add(url string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, exist := s.locations[url]; exist {
		return false
	}
	s.locations[url] = true
	return true
}
