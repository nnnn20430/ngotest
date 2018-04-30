// +build ignore

package main

import "fmt"
import "time"

func print_count_after_sec(i *int) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
	}
}

func main() {
	var i *int = new(int)
	go print_count_after_sec(i)
	// prevent optimization with too many layers of indirection
	_ = &i
	for {
		*i++
	}
}
