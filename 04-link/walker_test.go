package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

type expectedLink struct {
	href, text string
}

var tables = []struct {
	text          string
	expectedLinks []expectedLink
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
		[]expectedLink{
			{
				href: "/other-page",
				text: "A link to another page",
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
			t.Errorf("%s: len of parsed links %d not equal to expected len %d",
				testCaseString, len(parsedLinks), len(testCase.expectedLinks))
		}
		for j, link := range parsedLinks {
			expected := testCase.expectedLinks[j]
			if link.Href != expected.href {
				t.Errorf("%s: expected link %q != parsed %q", testCaseString, expected.href, link.Href)
			}
			if link.Text.String() != expected.text {
				t.Errorf("%s: expected text %q != parsed %q", testCaseString, expected.text, link.Text.String())
			}
		}
	}
}
