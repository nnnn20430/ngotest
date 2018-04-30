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


func noop() {
	// prevent inlining
	for {
		break
	}
	// slower alternative to prevent inlining
	//func(){}()
}

func main() {
	var i *int = new(int)
	go print_count_after_sec(i)
	for {
		*i++
		// prevent optimization
		noop()
	}
}
