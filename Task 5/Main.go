package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./Link"
)

func main() {
	HTMLfilePointer := flag.String("file", "htmlFormat.html",
		"File which contains paths and destinations in YAML format (default htmlFormat.html)")
	links, err := Link.ParseLinks(*HTMLfilePointer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(links)
	resp, err := http.Get("https://www.google.by")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
