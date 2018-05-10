// +build ignore

package main

import "fmt"
import "time"

func printCountAfterSec(i *uint) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
	}
}

func main() {
	var i = new(uint)
	go printCountAfterSec(i)
	for *i <= ^uint(0) {
		*i++
	}
}
