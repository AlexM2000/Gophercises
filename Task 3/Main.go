package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"./types"
)

func main() {
	JSONfilePointer := flag.String("File", "Story.json",
		"File which contains a story in JSON-file (default Story.json)")
	flag.Parse()
	adventureMap := parseJSON(*JSONfilePointer)
	h := NewHandler(adventureMap)
	log.Fatal(http.ListenAndServe(":8000", h))
}

var templ *template.Template

func NewHandler(adventureMap types.Story) http.Handler {
	return handler{adventureMap}
}

type handler struct {
	adventureMap types.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("htmlFormat.html")
	if err != nil {
		log.Fatal(err)
	}
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if charter, ok := h.adventureMap[path]; ok {
		templ.Execute(w, charter)
		return
	}

}

func parseJSON(filename string) map[string]types.Adventure {
	adventureMap := make(map[string]types.Adventure)
	JSONfile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(JSONfile)
	dec := json.NewDecoder(r)
	if err = dec.Decode(&adventureMap); err != nil {
		panic(err)
	}
	return adventureMap
}

// There should be html page
