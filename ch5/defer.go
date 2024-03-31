package main

import "fmt"

func main() {
	var f1, f2 = getFunc()
	f1("f1")
	defer f1("defer f1")
	defer f2("defer f2")
	f2("f2")
}

func getFunc() (func(s string), func(s string)) {
	return func(s string) {
			fmt.Println(s)
		}, func(s string) {
			defer fmt.Println("inner defer", s)
			fmt.Println(s)
		}
}
