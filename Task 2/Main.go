package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yaml-2"
)

func main() {
	yamlFilePointer := flag.String("yamlFile", "httpRequests.yaml",
		"File which contains paths and destinations in YAML format (default httpRequests.yaml)")
	flag.Parse()
	mux := http.NewServeMux()
	yamlHandler, err := YAMLHandler(*yamlFilePointer, mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Started serving port 8000")
	err = http.ListenAndServe(":8000", yamlHandler)
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
	PathUrls := make(map[string]string)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var PathUrl []PathAndUrl
	err = yaml.Unmarshal(yamlFile, &PathUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(PathUrl)
	for _, url := range PathUrl {
		PathUrls[url.Path] = url.URL
	}
	return PathUrls, err
}

type PathAndUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
