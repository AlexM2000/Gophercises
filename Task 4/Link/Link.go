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
	var link string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link = a.Val
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						links = append(links, Link{link, c.Data})
						break
					}
				}
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		iterNode(c)
	}
}
