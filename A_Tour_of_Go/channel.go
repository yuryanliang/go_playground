package main

import "fmt"

func sum(s []int, c chan int) {

	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("sum:", sum)

	c <- sum // send sum to c
}

func do(c chan int) {
	c <- 10
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x := <-c
	y := <-c
	//x, y := <-c, <-c // receive from c
	go do(c)
	fmt.Println(x, y, <-c)
}
