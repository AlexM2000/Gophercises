package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yaml-2"
)

func main() {
	yamlFile := flag.String("yamlFile", "httpRequests.yaml",
		"File which contains paths and destinations in YAML format (default httpRequests.yaml)")
	flag.Parse()
	mux := http.NewServeMux()
	yamlHandler, err := YAMLHandler(*yamlFile, mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Started serving port 8080")
	err = http.ListenAndServe(":8080", yamlHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusPermanentRedirect)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(filename)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathUrls, fallback), nil
}

func parseYaml(filename string) (map[string]string, error) {
	pathUrls := make(map[string]string)
	yamlFile, err := ioutil.ReadFile(filename)
	r := bytes.NewReader(yamlFile)
	dec := yaml.NewDecoder(r)
	var pathUrl pathAndUrl
	for dec.Decode(&pathUrl) == nil {
		pathUrls[pathUrl.Path] = pathUrl.URL
	}
	return pathUrls, err
}

type pathAndUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
