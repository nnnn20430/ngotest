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
	var s *bool = new(bool)
	go print_count_after_sec(i)
	for {
		*i++
		// prevent optimization
		if *s {
			break
		}
	}
}
