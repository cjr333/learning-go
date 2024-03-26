package main

import "fmt"

func main() {
	var x []int
	fmt.Println(len(x))
	y := append(x, 1, 2, 3)
	fmt.Println(x, len(x), cap(x))
	fmt.Println(y, len(y), cap(y))
	//y = append(y, 4)
	//fmt.Println(y, len(y), cap(y))
	y = append(y, 4, 5, 7, 8)
	fmt.Println(y, len(y), cap(y))
}
