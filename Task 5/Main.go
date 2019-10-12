package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"./Link"
)

type location struct {
	URL string `xml:"loc"`
}

type urlset struct {
	locs []location `xml:"url`
}

func HTMLWriteToFile(URL string, filename string) {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	HTMLfilePointer := flag.String("file", "htmlFormat.html",
		"File which contains paths and destinations (default htmlFormat.html)")
	URLpointer := flag.String("URL", "https://www.calhoun.io/",
		"URL, for which need to create sitemap (default https://www.calhoun.io/)")
	flag.Parse()
	GetSitemap(*URLpointer, *HTMLfilePointer)
}

func getLinksInPage(URL string, filename string) []Link.Link {
	HTMLWriteToFile(URL, filename)
	links, err := Link.ParseLinks(filename)
	if err != nil {
		log.Fatal(err)
	}
	return links
}

func getReqURL(URL string) *url.URL {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reqURL := resp.Request.URL
	return reqURL
}

func getFilteredLinksHref(links []Link.Link, baseURL string) []string {
	var linksHref, linksFilteredHref []string
	for _, v := range links {
		switch {
		case strings.HasPrefix(v.Href, "/"):
			linksHref = append(linksHref, baseURL+v.Href)
		case strings.HasPrefix(v.Href, "http"):
			linksHref = append(linksHref, v.Href)
		}
	}
	//fmt.Println(linksHref)
	for _, v := range linksHref {
		switch {
		case strings.HasPrefix(v, baseURL):
			linksFilteredHref = append(linksFilteredHref, v)
		}
	}
	//fmt.Println(linksFilteredHref)
	return linksFilteredHref
}

func GetSitemap(URL string, filename string) {
	visited := make(map[string]bool)
	appended := make(map[string]bool)
	links := getLinksInPage(URL, filename)
	reqURL := getReqURL(URL)
	fmt.Println(reqURL)
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	baseURLstr := baseURL.String()
	fmt.Println(baseURLstr)
	linksFilteredHref := getFilteredLinksHref(links, baseURLstr)
	fmt.Println(linksFilteredHref)
	for _, v := range linksFilteredHref {
		if _, ok := visited[v]; !ok {
			tempLinks := getLinksInPage(v, filename)
			tempFilteredLinks := getFilteredLinksHref(tempLinks, baseURLstr)
			for _, v2 := range tempFilteredLinks {
				if _, ok := appended[v2]; !ok {
					linksFilteredHref = append(linksFilteredHref, v2)
					appended[v2] = true
				}
			}
			visited[v] = true
		}
	}
	fmt.Println(linksFilteredHref)
}
