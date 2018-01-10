package main

import "fmt"
import "./cgohello"

func main() {
	fmt.Printf("simple call:\n")
	cgohello.World()
	fmt.Printf("\nuse GoString():\n")
	cgohello.PtrToStr()
	fmt.Printf("\nconvert pointer to array:\n")
	cgohello.PtrToArr()
	fmt.Printf("\nuse pointer arithmetic:\n")
	cgohello.PtrArithmetic()
}
