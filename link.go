package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents html anchor tag with href attribute and text inside
type Link struct {
	Href string
	Text string
}

// ExtractLinks searches valid html for anchors tags
func ExtractLinks(r io.Reader) ([]Link, error) {
	document, error := html.Parse(r)
	if error != nil {
		return nil, error
	}

	links := parseForLinks(document)
	return links, nil
}

func parseForLinks(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		links = append(links, Link{
			Href: extractAttributeValue(n, "href"),
			Text: strings.TrimSpace(strings.Trim(extractLinkText(n), "\n")),
		})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseForLinks(c)...)
	}
	return links
}

func extractAttributeValue(n *html.Node, attr string) string {
	for _, v := range n.Attr {
		if v.Key == attr {
			return v.Val
		}
	}
	return ""
}

func extractLinkText(n *html.Node) string {
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.ElementNode:
			text = fmt.Sprint(text, extractLinkText(c))
		case html.TextNode:
			text = text + c.Data
		}
	}
	return text
}
