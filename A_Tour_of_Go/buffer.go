package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	go func() { ch <- 4 }() //sol: add the ch<-4 in a goroutine
	//ch<-4//cause deadlock,
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
