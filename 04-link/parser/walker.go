package parser

import (
	"golang.org/x/net/html"
	"log"
	"strings"
)

type Link struct {
	Href, Text string
}

type walker struct {
	Links []Link
}

func NewWalker(text string) *walker {
	r := strings.NewReader(text)
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	w := &walker{}
	w.walk(doc)
	return w
}

func (w *walker) walk(node *html.Node) {
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
		w.walk(child)
	}
}

func (w *walker) getText(node *html.Node, builder *strings.Builder) {
	if node.Type == html.TextNode {
		builder.WriteString(node.Data)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		w.getText(child, builder)
	}

}
