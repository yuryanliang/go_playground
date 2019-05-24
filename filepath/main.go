package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir("t"))
	fmt.Println(dir)

}
