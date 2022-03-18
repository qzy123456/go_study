package main

import (
	"fmt"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {

	http.Handle("/foo", loggingMiddleware(http.HandlerFunc(foo)))
	http.Handle("/bar", loggingMiddleware(http.HandlerFunc(bar)))
	http.ListenAndServe(":8080", nil)
}
