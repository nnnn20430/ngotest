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

func noop() {
	// prevent inlining
	for {
		break
	}
	// slower alternative to prevent inlining
	//func(){}()
}

func main() {
	var i = new(int)
	go printCountAfterSec(i)
	for {
		*i++
		// prevent optimization
		noop()
	}
}
