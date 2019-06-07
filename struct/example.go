package main

import "fmt"

type a struct {
	name string
}

type b struct {
	a
}

type inter interface {
	DoSmth()
}

func (x a) DoSmth() {
	fmt.Println(x.name)
}

func main() {
	var slice []inter
	for i := 0; i < 5; i++ {

		slice = append(slice, &b{a: a{name: "Tom"}})
	}

	for _, a := range slice {
		println("a in slice", &a)
	}

	// Doesn't work
	for i, _ := range slice {
		x := slice[i].(*b)
		x.name = "Tim"
	}

	// Compiler error, the interface doesn't have a .name
	//for i, _ := range slice {
	//	slice[i].name = "Tim"
	//}

	for _, x := range slice {
		x.DoSmth()
	}
}
