package main

import "fmt"

func asmGetStr() string

func main() {
	fmt.Printf("%v", asmGetStr())
}
