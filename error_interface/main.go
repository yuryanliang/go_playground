package main

import (
	"fmt"
	"github.com/yuryanliang/go_playground/error_interface/database"
)

func main() {
	var id = 7
	var c = database.CustomStruct{Num: id}
	fmt.Println("c.Find:", c.Find(id))

	newRepoString := database.NewRepositoryString(1)
	fmt.Println(
		"newRepoString.Find:", newRepoString.Find(900),
	)

}
