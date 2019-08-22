package main

import (
	"fmt"
	"strconv"
	"time"
)

func er() {
	i, err := strconv.Atoi("42s")
	if err != nil {
		fmt.Printf("couldn't convert number:%T, %v\n", err, err)
		return
	}
	fmt.Println(i)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(), "not work",
	}
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
