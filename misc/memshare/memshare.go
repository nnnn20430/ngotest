// +build ignore

package main

import "fmt"
import "time"

func resCountAfterSec(ir *int) {
	for {
		time.Sleep(1 * time.Second)
		*ir = 1
	}
}

func main() {
	var i, ir int
	go resCountAfterSec(&ir)
	for {
		i++
		if ir == 1 {
			fmt.Printf("%d\n", i)
			i, ir = 0, 0
		}
	}
}
