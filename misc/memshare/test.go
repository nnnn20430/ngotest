// +build ignore

package main

import "fmt"
import "time"

func print_count_after_sec(i *uint) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
	}
}

func main() {
	var i *uint = new(uint)
	go print_count_after_sec(i)
	for *i <= ^uint(0) {
		*i++
	}
}
