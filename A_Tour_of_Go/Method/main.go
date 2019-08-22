package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X * v.Y
}

func (v Vertex) Sub() float64 {
	return v.X - v.Y
}

func me() {
	v := Vertex{2, 3}
	res := v.Sub()
	fmt.Println(res)
}

func main() {
	me()
}
