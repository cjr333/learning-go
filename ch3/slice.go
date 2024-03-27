package main

import "fmt"

func main() {
	x := append([]int{}, 1, 2, 3)
	fmt.Println(x, len(x), cap(x))

	y := append(x, 4)
	fmt.Println(y, len(y), cap(y))

	z := append(x, 4, 5, 7, 8)
	fmt.Println(z, len(z), cap(z))

	i := append(y, 4, 5, 7, 8)
	fmt.Println(i, len(i), cap(i))
}
