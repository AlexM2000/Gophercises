package main

import (
	"flag"
	"fmt"
	"log"

	"Gophercises/Task 4/Link"
)

func main() {
	HTMLfilePointer := flag.String("file", "htmlFormat.html",
		"File which contains paths and destinations in YAML format (default htmlFormat.html)")
	links, err := Link.ParseLinks(*HTMLfilePointer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(links)
}
