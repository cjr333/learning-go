package main

import "fmt"

func main() {
	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	fp := firstPerson{
		name: "Bob",
		age:  50,
	}

	sp := secondPerson(fp)
	fmt.Println(sp)

	var g struct {
		name string
		age  int
	}
	//g = fp
	g.name = "Bob"
	g.age = 50
	fmt.Println(fp == g)
}
