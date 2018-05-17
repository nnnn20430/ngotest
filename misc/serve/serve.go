package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Printf("http://localhost:8080/\n")
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
