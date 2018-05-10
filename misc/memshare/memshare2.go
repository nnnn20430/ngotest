// +build ignore

package main

import "fmt"
import "time"

func resCountAfterSec(i, ir *int) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
		*ir = 1
	}
}

func main() {
	var i, ir int
	go resCountAfterSec(&i, &ir)
	for {
		i++
		if ir == 1 {
			i, ir = 0, 0
		}
	}
}
