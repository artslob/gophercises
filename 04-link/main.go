package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

type Link struct {
	Href, Text string
}

type Walker struct {
	Links []Link
}

func (w *Walker) Walk(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		newLink := Link{}
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				newLink.Href = attr.Val
				break
			}
		}
		var builder strings.Builder
		w.getText(node, &builder)
		newLink.Text = strings.TrimSpace(builder.String())
		w.Links = append(w.Links, newLink)
		return
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		w.Walk(child)
	}
}

func (w *Walker) getText(node *html.Node, builder *strings.Builder) {
	if node.Type == html.TextNode {
		builder.WriteString(node.Data)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		w.getText(child, builder)
	}

}

func main() {
	text := `
		<p>Links:</p>
		<ul>
			<li><a href="/first">Foo <b>bold</b></a></li>
			<li><a href="/second">BarBaz <span>in span <b>bold<b/></span> <b>test</b></a></li>
		</ul>
	`
	r := strings.NewReader(text)
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	walker := Walker{}
	walker.Walk(doc)
	for _, link := range walker.Links {
		fmt.Println("link:", link.Href)
		fmt.Println("text:", link.Text)
	}
}
