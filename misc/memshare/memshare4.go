package main

import "fmt"
import "time"

func res_count_after_sec(c chan int) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d\n", <-c)
	}
}

func method_one(c chan int) {
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

func method_two(c chan int) {
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
	go res_count_after_sec(c)
	method_two(c)
}
