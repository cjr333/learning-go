package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	//done := false
	//for !done {
	//	select {
	//	case ch2 <- v:
	//		fmt.Println(v, v2)
	//		done = true
	//	case v2 = <-ch1:
	//		fmt.Println("read from ch1")
	//	default:
	//		time.Sleep(500)
	//	}
	//}
	//time.Sleep(1000)
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
	time.Sleep(2000)
}
