package linkparser

import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"regexp"
	"strings"
)

type htmlLink struct {
	href, text string
}

func (link *htmlLink) Href() string {
	return link.href
}

func (link *htmlLink) Text() string {
	return link.text
}

func SimpleLinkParser() HtmlLinkParser {
	return HtmlLinkParserFunc(parseLinks)
}

func parseLinks(htmlText []byte) ([]HtmlLink, error) {
	reader := bytes.NewReader(htmlText)
	node, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return crawl(node, make([]HtmlLink, 0)), nil
}

func crawl(node *html.Node, links []HtmlLink) []HtmlLink {
	if node.Type == html.ElementNode && node.DataAtom == atom.A {
		link, err := linkFromNode(node)
		if err != nil {
			panic(err)
		}
		links = append(links, link)
	}
	for next := node.FirstChild; next != nil; next = next.NextSibling {
		links = crawl(next, links)
	}
	return links
}

func linkFromNode(node *html.Node) (HtmlLink, error) {
	if node.Type != html.ElementNode {
		return nil, errors.New("node must be an ElementNode")
	}
	if node.DataAtom != atom.A {
		return nil, errors.New("element must be an anchor")
	}

	result := new(htmlLink)
	for _, attr := range node.Attr {
		if strings.ToLower(attr.Key) == "href" {
			result.href = attr.Val
		}
	}
	result.text = buildLinkText(node)

	return result, nil
}

var whitespaceRegexp = regexp.MustCompile(`(\s+)`)

func buildLinkText(node *html.Node) string {
	builder := new(strings.Builder)
	recBuildLinkText(node, builder)
	result := builder.String()
	result = strings.TrimSpace(result)
	result = whitespaceRegexp.ReplaceAllString(result, " ")
	return result
}

func recBuildLinkText(node *html.Node, builder *strings.Builder) {
	switch node.Type {
	case html.TextNode:
		builder.WriteString(node.Data)
	case html.ElementNode:
		for next := node.FirstChild; next != nil; next = next.NextSibling {
			recBuildLinkText(next, builder)
		}
	}
}
