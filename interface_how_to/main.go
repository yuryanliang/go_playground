package main

import (
	"fmt"
	. "github.com/yuryanliang/go_playground/interface_how_to/type"
)

func main() {
	animals := []Animal{
		Dog{},
		Cat{},
		Llama{},
		JavaProgrammer{},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	fmt.Println(Eat(Apple{}))
	fmt.Println(Eat("sting"))
	fmt.Println(Eat(2))

}
