package main

import (
	"fmt"
	"strings"
)

func ar() {
	var a [10]int
	fmt.Println(a)
	a[0] = 1
	fmt.Println(a)

	b := [3]int{1, 3, 2}
	fmt.Println(b[0])
	s := b[1:3]
	fmt.Println(s)

}
func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v \n", len(s), cap(s), s)
}
func sl() {
	s := []int{2, 3, 4, 6, 7, 12}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:6]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	var x []int
	y := make([]int, 0, 5)
	printSlice(y)

	if x == nil {
		fmt.Print("nil")
	}

}

func tic() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	board[0][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}

}

func ap() {
	var s []int
	s = append(s, 0)
	printSlice(s)

	s = append(s, 3, 3, 4)
	printSlice(s)
}

func ma() {
	m := map[string]int{}
	//var m1 map[string]int
	var m2 = map[string]int{
		"a": 1,
		"b": 2,
	}
	m["a"] = 1
	fmt.Println(m2)

	delete(m2, "a")
	fmt.Println(m2)

	elem, ok := m2["b"]
	fmt.Println(elem, ok)

}
func main() {
	//ar()
	//sl()
	//tic()
	//ap()
	ma()
}
