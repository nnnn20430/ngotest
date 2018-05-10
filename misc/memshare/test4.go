// +build ignore

package main

import "fmt"
import "time"

func printCountAfterSec(i *int) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
	}
}

func main() {
	var i = new(int)
	go printCountAfterSec(i)
	// prevent optimization with too many layers of indirection
	_ = &i
	for {
		*i++
	}
}
