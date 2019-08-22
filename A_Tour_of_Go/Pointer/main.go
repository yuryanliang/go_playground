package main

import "fmt"

func first() {
	//var p *int
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
}

//A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func struc() {
	v := Vertex{1, 2}

	var p *Vertex
	p = &v

	v1 := Vertex{1, 2}
	v2 := Vertex{X: 3, Y: 4}
	v3 := Vertex{}
	p1 := &Vertex{1, 2}
	fmt.Println(v1, v2, v3, p1)
	fmt.Printf("%v", p1)

	fmt.Printf("type %T, value %v", *p, p.X)
}

func main() {
	//first()
	struc()
}
