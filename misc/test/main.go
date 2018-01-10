package main

import "fmt"

var tm []byte = make([]byte, 10)
var test *[]byte = &tm

func lol(thing *[]byte) {
	(*thing)[0] = 10
}

func main() {
	var eh []byte = (*test)[:len(*test)]
	var b []byte = append([]byte(nil), eh...)
	lol(&b)
	fmt.Printf(string((*test)[0]))
	tm = nil
	eh = nil
	b = nil
}
