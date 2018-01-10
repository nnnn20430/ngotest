package main

import "fmt"
import . "hello"

func main() {
	fmt.Printf("simple call:\n")
	Hello()
	fmt.Printf("\nuse GoString():\n")
	PtrToStr()
	fmt.Printf("\nconvert pointer to array:\n")
	PtrToArr()
	fmt.Printf("\nuse pointer arithmetic:\n")
	PtrArithmetic()
}
