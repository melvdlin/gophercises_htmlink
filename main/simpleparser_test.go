package main

import (
	linkparser2 "htmlink/linkparser"
	utils2 "htmlink/utils"
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
	example1Expected := []linkparser2.HtmlLink{
		&expectedLink{
			href: "/other-page",
			text: "A link to another page",
		},
	}
	example2Expected := []linkparser2.HtmlLink{
		&expectedLink{
			href: "https://www.twitter.com/joncalhoun",
			text: "Check me out on twitter",
		},
		&expectedLink{
			href: "https://github.com/gophercises",
			text: "Gophercises is on Github!",
		},
	}
	example3Expected := []linkparser2.HtmlLink{
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
	example4Expected := []linkparser2.HtmlLink{
		&expectedLink{
			href: "/dog-cat",
			text: "cog cat",
		},
	}

	parserUnderTest = linkparser2.SimpleLinkParser()

	expectedLinks = example1Expected
	htmlText, err = utils2.ReadFile(example1Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex1", testParser)

	expectedLinks = example2Expected
	htmlText, err = utils2.ReadFile(example2Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex2", testParser)

	expectedLinks = example3Expected
	htmlText, err = utils2.ReadFile(example3Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex3", testParser)

	expectedLinks = example4Expected
	htmlText, err = utils2.ReadFile(example4Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Ex4", testParser)
}

var parserUnderTest linkparser2.HtmlLinkParser
var expectedLinks []linkparser2.HtmlLink
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
		if !utils2.AnyInSlice(links, func(_ int, link linkparser2.HtmlLink) bool {
			return link.Href() == expected.Href() && link.Text() == expected.Text()
		}) {
			t.Fail()
		}
	}
}
