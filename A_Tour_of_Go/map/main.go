package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	fmt.Println("split", strings.Fields(s))
	l := strings.Fields(s)
	res := make(map[string]int)
	for _, v := range l {

		if _, ok := res[v]; ok {
			res[v] += 1
		} else {
			res[v] = 1
		}
	}
	return res
}

func main() {
	s := "I ate a donut. Then I ate another donut."
	fmt.Println(WordCount(s))
}
