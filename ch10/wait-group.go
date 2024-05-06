package main

import (
	"fmt"
	"sync"
)

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			v := <-in
			//for v := range in {
			out <- processor(v)
			//}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

func main() {
	num := 5
	ch := make(chan int)
	double := func(x int) int {
		return x * 2
	}
	go func() {
		for i := 0; i < num; i++ {
			ch <- i
		}
	}()
	result := processAndGather(ch, double, num)
	fmt.Println(result)

}
