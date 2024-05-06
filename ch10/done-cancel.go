package main

import (
	"fmt"
	"time"
)

func countTo2(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				fmt.Println("done")
				return
			case ch <- i:
				//default:
				//	ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}
func main() {
	ch, cancel := countTo2(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
	time.Sleep(2 * time.Second)
}
