package Link

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var links []Link

func ParseLinks(HTMLfilename string) ([]Link, error) {
	htmlBytes, err := ioutil.ReadFile(HTMLfilename)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(htmlBytes)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	iterNode(doc)
	return links, nil
}

func iterNode(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		iterateOnTagA(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		iterNode(c)
	}
}

func iterateOnTagA(n *html.Node) {
	for _, a := range n.Attr {
		iterateOnHref(a, n)
	}
}

func iterateOnHref(a html.Attribute, n *html.Node) {
	if a.Key == "href" {
		link := a.Val
		appendTextOfLink(a, n, link)
	}
}

func appendTextOfLink(a html.Attribute, n *html.Node, link string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			links = append(links, Link{link, c.Data})
			break
		}
	}
}
