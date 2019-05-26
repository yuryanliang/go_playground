package main

import "fmt"

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"a", "b", "c"}
	vals := make([]interface{}, len(names))
	for i, n := range names {
		vals[i] = n
	}

	//PrintAll(names)
	PrintAll(vals)
}
