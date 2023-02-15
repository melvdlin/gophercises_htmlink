package linkparser

type HtmlLink interface {
	Href() string
	Text() string
}

type HtmlLinkParser interface {
	Parse(htmlText []byte) ([]HtmlLink, error)
}

type HtmlLinkParserFunc func([]byte) ([]HtmlLink, error)

func (f HtmlLinkParserFunc) Parse(htmlText []byte) ([]HtmlLink, error) {
	return f(htmlText)
}
