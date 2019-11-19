package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/panic/", panicDemo)
	r.HandleFunc("/panic-after/", panicAfterDemo)
	r.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8000", handlerWithPanic(r, false)))
}

func handlerWithPanic(h http.Handler, devMode bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if devMode {
					log.Println(string(debug.Stack()))
					fmt.Fprintln(w, string(debug.Stack()))
					return
				}
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
