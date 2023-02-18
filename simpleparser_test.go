package main

import (
	"github.com/melvdlin/gophercises_htmlink/linkparser"
	"github.com/melvdlin/gophercises_htmlink/utils"
	"path/filepath"
	"testing"
)

func TestParser(t *testing.T) {
	example1Path, err := filepath.Abs("resources/ex1.html")
	example2Path, err := filepath.Abs("resources/ex2.html")
	example3Path, err := filepath.Abs("resources/ex3.html")
	example4Path, err := filepath.Abs("resources/ex4.html")
	if err != nil {
		t.Fatal(err)
	}
	example1Expected := []linkparser.HtmlLink{
		&expectedLink{
			href: "/other-page",
			text: "A link to another page",
		},
	}
	example2Expected := []linkparser.HtmlLink{
		&expectedLink{
			href: "https://www.twitter.com/joncalhoun",
			text: "Check me out on twitter",
		},
		&expectedLink{
			href: "https://github.com/gophercises",
			text: "Gophercises is on Github!",
		},
	}
	example3Expected := []linkparser.HtmlLink{
		&expectedLink{
			href: "#",
			text: "Login",
		},
		&expectedLink{
			href: "/lost",
			text: "Lost? Need help?",
		},
		&expectedLink{
			href: "https://twitter.com/marcusolsson",
			text: "@marcusolsson",
		},
	}
	example4Expected := []linkparser.HtmlLink{
		&expectedLink{
			href: "/dog-cat",
			text: "cog cat",
		},
	}

	parserUnderTest = linkparser.SimpleLinkParser()

	expectedLinks = example1Expected
	htmlText, err = utils.ReadFile(example1Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex1", testParser)

	expectedLinks = example2Expected
	htmlText, err = utils.ReadFile(example2Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex2", testParser)

	expectedLinks = example3Expected
	htmlText, err = utils.ReadFile(example3Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex3", testParser)

	expectedLinks = example4Expected
	htmlText, err = utils.ReadFile(example4Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex4", testParser)
}

var parserUnderTest linkparser.HtmlLinkParser
var expectedLinks []linkparser.HtmlLink
var htmlText []byte

type expectedLink struct {
	href, text string
}

func (link *expectedLink) Href() string {
	return link.href
}

func (link *expectedLink) Text() string {
	return link.text
}

func testParser(t *testing.T) {

	links, err := parserUnderTest.Parse(htmlText)
	if err != nil {
		t.Fatal(err)
	}
	if len(links) != len(expectedLinks) {
		t.Fail()
	}
	for _, expected := range expectedLinks {
		if !utils.AnyInSlice(links, func(_ int, link linkparser.HtmlLink) bool {
			return link.Href() == expected.Href() && link.Text() == expected.Text()
		}) {
			t.Fail()
		}
	}
}
