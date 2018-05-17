// +build ignore

package main

import (
	"net/http"
	"fmt"
)

type htmlString string

func (s htmlString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", s)
}

func main() {
	http.Handle("/", htmlString("Hello, world."))
	fmt.Printf("http://localhost:8080/\n")
	panic(http.ListenAndServe(":8080", nil))
}
