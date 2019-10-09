package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./Link"
)

func HTMLWriteToFile(url string, filename string) {
	resp, err := http.Get(url)
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
		"File which contains paths and destinations in YAML format (default htmlFormat.html)")
	_ = flag.String("URL", "https://www.calhoun.io/",
		"URL, for which need to create sitemap (default https://www.calhoun.io/)")
	flag.Parse()
	HTMLWriteToFile("https://www.calhoun.io/", *HTMLfilePointer)
	links, err := Link.ParseLinks(*HTMLfilePointer)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range links {
		fmt.Println("Href: ", v.Href)
		fmt.Println("Text: ", v.Text)
	}
}
