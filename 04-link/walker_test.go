package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

var tables = []struct {
	text          string
	expectedLinks []Link
}{
	{
		`
		<html>
		<body>
			<h1>Hello!</h1>
			<a href="/other-page">A link to another page</a>
		</body>
		</html>
		`,
		[]Link{
			{
				Href: "/other-page",
				Text: "A link to another page",
			},
		},
	},
}

func TestWalker(t *testing.T) {
	for i, testCase := range tables {
		testCaseString := fmt.Sprintf("test case %d", i)
		r := strings.NewReader(testCase.text)
		doc, _ := html.Parse(r)
		walker := Walker{}
		walker.Walk(doc)
		parsedLinks := walker.Links
		if len(parsedLinks) != len(testCase.expectedLinks) {
			t.Fatalf("%s: len of parsed links %d not equal to expected len %d",
				testCaseString, len(parsedLinks), len(testCase.expectedLinks))
		}
		for j, link := range parsedLinks {
			expected := testCase.expectedLinks[j]
			if link.Href != expected.Href {
				t.Errorf("%s: expected link %q != parsed %q", testCaseString, expected.Href, link.Href)
			}
			if link.Text != expected.Text {
				t.Errorf("%s: expected text %q != parsed %q", testCaseString, expected.Text, link.Text)
			}
		}
	}
}
