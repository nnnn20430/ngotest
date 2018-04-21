package main

import "fmt"

func argslice(args ...interface{}) []interface{} {
	return args
}

func div(i1, i2 int) (int, int) {
	return i1/i2, i1%i2
}

func main() {
	fmt.Printf("5/2=%d, 5%%2=%d\n", argslice(div(5, 2))[0], argslice(div(5, 2))[1])
}
