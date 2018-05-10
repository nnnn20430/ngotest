package main

import "fmt"
import "time"

func resCountAfterSec(c chan int) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", <-c)
	}
}

func methodOne(c chan int) {
	i := int(0)
	for {
		i++
		select {
		case c <- i:
			i = 0
		default:
		}
	}
}

func methodTwo(c chan int) {
	i := int(0)
	for {
		for ir := 0; ir < 1000; ir++ {
			i++
		}
		select {
		case c <- i:
			i = 0
		default:
		}
	}
}

func main() {
	c := make(chan int)
	go resCountAfterSec(c)
	methodTwo(c)
}
