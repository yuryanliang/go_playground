package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

func sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for i < 10 {
		z -= (z*z - x) / (2 * z)
		fmt.Println("for loop:", i, z, z*z)
		i += 1
	}
	return z
}

func sw() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		//freebsd, openbsd
		//plan9
		fmt.Printf("%s. \n", os)
	}
}

func de() {
	i := 0
	for i < 10 {
		defer fmt.Println(i)
		i += 1
	}
}

func ra() {
	pow := []int{1, 2, 4}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
func main() {
	ra()
	//de()
	//defer fmt.Println("Defer")
	//sw()
	//z:=sqrt(50.0)
	//fmt.Println(z*z)
	//sum := 1
	//for sum <10{
	//	sum +=sum
	//	fmt.Println( "sum", sum)
	//
	//}
	//fmt.Println(sum)
	//
	//fmt.Println(pow(3, 2, 10))
	//fmt.Println(pow(3, 3, 10))

}
