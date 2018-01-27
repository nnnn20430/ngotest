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
	http.ListenAndServe(":8000", nil)
}
