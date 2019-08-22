package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum - 2
	y = x - 2
	return
}

//var c, python, java bool
//var i, j int = 1, 2

//var(
//	ToBe bool = false
//	MaxInt uint64 = 1<<64 -1
//	z complex128 = cmplx.Sqrt(-5 + 12i)
//)

//const World = "world"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))

	//fmt.Println("My favorite number is", rand.Intn(6), math.Sqrt(7))
	//fmt.Printf("type %g problem. \n", math.Sqrt(8))
	//fmt.Printf("type %g problem. \n", 6)
	//fmt.Println(math.Pi)
	//fmt.Println("result ", add(1, 2))
	//a, b := swap("hello", "world")
	//fmt.Println("result ", a, b)
	//a, b := split(10)
	//fmt.Println("result ", a, b)
	//var ii int
	//fmt.Println(i,ii, j)

	//c, python, java := true, false, "no"
	//fmt.Println(c, python, java)
	//
	//k :=3
	//fmt.Println("k", k)
	//fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	//fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	//fmt.Printf("%g,Type: %T Value: %v \n", z,z, z)
	//var i int = 42
	//var f float64 = float64(i)
	//var u uint = uint(f)

	//var x, y int = 3, 4
	//var f  = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)
	//fmt.Println(x, f, z)

	//v := 0.8 + 5i // change me!
	//fmt.Printf("v is of type %T\n", v)

	//const P = 4.3
	//fmt.Println(P, World)

}
