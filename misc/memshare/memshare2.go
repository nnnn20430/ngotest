// +build ignore

package main

import "fmt"
import "time"

func res_count_after_sec(i, ir *int) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
		*ir = 1
	}
}

func main() {
	var i, ir int
	go res_count_after_sec(&i, &ir)
	for {
		i++
		if ir == 1 {
			i, ir = 0, 0
		}
	}
}
