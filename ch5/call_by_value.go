package main

import "fmt"

func main() {
	var m = map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	var s = []int{1, 2, 3}
	fmt.Println(&s)
	fmt.Println(&s[0])
	modSlice(s)

	var x = 1
	fmt.Println(&x)
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	fmt.Println(&s[0])
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
	fmt.Println(&s[0])
}
