// +build ignore

package main

import "fmt"
import "time"
import "sync"

var l sync.Mutex

func res_count_after_sec(i *int) {
	for {
		l.Lock()
		*i = 0
		l.Unlock()
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
	}
}

func main() {
	var i *int = new(int)
	go res_count_after_sec(i)
	for {
		l.Lock()
		*i++
		l.Unlock()
	}
}
