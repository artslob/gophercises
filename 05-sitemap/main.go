package main

import (
	"flag"
	"fmt"
	"github.com/artslob/gophercises/05-sitemap/sitemap"
)

func main() {
	url := flag.String("url", "google.com", "links map of this url is built")
	s := sitemap.NewSiteMap()
	s.BuildMap(*url)
	for link := range s.GetMap() {
		fmt.Println(link)
	}
}
