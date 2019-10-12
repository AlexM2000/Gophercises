package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"./Link"
)

type location struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Locs []location `xml:"url"`
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

func outputToXMLfile(links []string, filename string) {
	URLS := urlset{}
	for _, link := range links {
		URLS.Locs = append(URLS.Locs, location{link})
	}
	file, _ := os.Create(filename)
	xmlWriter := io.Writer(file)
	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(URLS); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	HTMLfilePointer := flag.String("file", "htmlFormat.html",
		"File which contains paths and destinations (default htmlFormat.html)")
	URLpointer := flag.String("URL", "http://www.voprospsyha.narod.ru/", //Сайт нашего преподавателя по психологии хехе
		"URL, for which need to create sitemap (default http://www.sitemaps.org/schemas/sitemap/0.9)")
	flag.Parse()
	links := GetSitemap(*URLpointer, *HTMLfilePointer)
	outputToXMLfile(links, "Output.xml")
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

func GetSitemap(URL string, filename string) []string {
	visited := make(map[string]bool)
	appended := make(map[string]bool)
	links := getLinksInPage(URL, filename)
	reqURL := getReqURL(URL)
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	baseURLstr := baseURL.String()
	linksFilteredHref := getFilteredLinksHref(links, baseURLstr)
	length := len(linksFilteredHref)
	for i := 0; i < length; i++ {
		fmt.Println(linksFilteredHref)
		if _, ok := visited[linksFilteredHref[i]]; !ok {
			visited[linksFilteredHref[i]] = true
			tempLinks := getLinksInPage(linksFilteredHref[i], filename)
			tempFilteredLinks := getFilteredLinksHref(tempLinks, baseURLstr)
			fmt.Println(tempFilteredLinks)
			for _, v2 := range tempFilteredLinks {
				if _, ok := appended[v2]; !ok {
					linksFilteredHref = append(linksFilteredHref, v2)
					appended[v2] = true
				}
			}
		}
		length = len(linksFilteredHref)
	}
	return linksFilteredHref
}
