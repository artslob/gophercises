package main

import (
	"fmt"
	"github.com/artslob/gophercises/04-link/parser"
)

func main() {
	text := `
		<p>Links:</p>
		<ul>
			<li><a href="/first">Foo <b>bold</b></a></li>
			<li><a href="/second">BarBaz <span>in span <b>bold<b/></span> <b>test</b></a></li>
		</ul>
	`
	walker := parser.NewWalker(text)
	for _, link := range walker.Links {
		fmt.Println("link:", link.Href)
		fmt.Println("text:", link.Text)
	}
}
