package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var links []Link

func main() {
	HTMLfilePointer := flag.String("file", "htmlFormat.html",
		"File which contains paths and destinations in YAML format (default htmlFormat.html)")
	htmlBytes, err := ioutil.ReadFile(*HTMLfilePointer)
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader(htmlBytes)
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	iterNode(doc)
	file, err := json.MarshalIndent(links, " ", " ")
	err = ioutil.WriteFile("Links.json", file, 0644)
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
