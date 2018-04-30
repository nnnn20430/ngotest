// +build ignore

package main

import "fmt"
import "time"

func res_count_after_sec(ir *int) {
	for {
		time.Sleep(1 * time.Second)
		*ir = 1
	}
}

func main() {
	var i, ir int
	go res_count_after_sec(&ir)
	for {
		i++
		if ir == 1 {
			fmt.Printf("%d\n", i)
			i, ir = 0, 0
		}
	}
}
