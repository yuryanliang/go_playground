package main

import "fmt"

func sum(nums []int, c chan int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	c <- sum
}

func routine() {
	nums := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(nums[:3], c)
	go sum(nums[3:], c)
	a := <-c
	b := <-c
	fmt.Println(a, b, a+b)

}

func main() {
	routine()
}
