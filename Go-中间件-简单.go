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

func use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {

	http.Handle("/foo", use(http.HandlerFunc(foo)))
	http.Handle("/bar", use(http.HandlerFunc(bar)))
	http.ListenAndServe(":8080", nil)
}
