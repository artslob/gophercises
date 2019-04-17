package main

import (
	"flag"
	"fmt"
	"github.com/artslob/gophercises/05-sitemap/sitemap"
	"log"
)

func main() {
	url := flag.String("url", "google.com", "links map of this url is built")
	s := sitemap.NewSiteMap()
	s.BuildMap(*url)
	bytes, err := sitemap.GetXml(s.GetMap())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bytes)
}
