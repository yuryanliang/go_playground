package main

import (
	"fmt"
	"sync"
	"time"
)

func SomeFunc(wg *sync.WaitGroup) {

	fmt.Println("goroutine Start")
	time.Sleep(time.Second)
	fmt.Println("goroutine done")
	wg.Done()
}

func main() {
	//wg := sync.WaitGroup{}
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("Progrem Start")
	go SomeFunc(&wg)
	wg.Wait()
	fmt.Println("Program done")
}
