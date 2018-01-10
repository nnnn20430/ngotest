package main

import "fmt"

func main() {
	var test string = string(byte(uint8(rune('"'))))
	fmt.Println([]byte(test)[0])
}
