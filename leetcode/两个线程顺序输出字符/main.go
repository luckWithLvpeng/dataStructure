package main

import (
	"fmt"
	"sync"
)

func main() {
	var s1 = "aceg"
	var s2 = "bdfh"
	var c1 = make(chan int)
	var c2 = make(chan int)
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func(c1 chan int, c2 chan int) {
		i := 0
		for true {
			select {
			case <-c1:
				fmt.Printf("%c", s1[i])
				i++
				c2 <- 1
				if i == len(s1) {
					wg.Done()
					return
				}
			}

		}
	}(c1, c2)
	go func(c1 chan int, c2 chan int) {
		i := 0
		for true {
			select {
			case <-c2:
				fmt.Printf("%c", s2[i])
				i++
				if i == len(s2) {
					wg.Done()
					return
				}
				c1 <- 1
			}

		}
	}(c1, c2)
	c1 <- 1
	wg.Wait()
	fmt.Println()
	fmt.Println("结束")
}
