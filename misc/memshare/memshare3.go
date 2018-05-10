// +build ignore

package main

import "fmt"
import "time"
import "sync"

var l sync.Mutex

func resCountAfterSec(i *int) {
	for {
		l.Lock()
		*i = 0
		l.Unlock()
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", *i)
	}
}

func main() {
	var i = new(int)
	go resCountAfterSec(i)
	for {
		l.Lock()
		*i++
		l.Unlock()
	}
}
