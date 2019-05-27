package main

import "fmt"

type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type db struct {
}

func (d db) FetchMessage(lang string) (string, error) {
	return "Fetch Message call :" + lang, nil
}
func (d db) FetchDefaultMessage() (string, error) {
	return "Default Message call ", nil
}

type GreeterService interface {
	Greet() string
	GreetDefaultMsg() string
}

type greeter struct {
	database db
	lang     string
}

func (g greeter) Greet() string {
	word, _ := g.database.FetchMessage(g.lang)
	return word
}

func (g greeter) GreetDefaultMsg() string {
	word, _ := g.database.FetchDefaultMessage()
	return word
}

func main() {
	actualDatabase := db{}
	greeter := greeter{
		database: actualDatabase,
		lang:     "English",
	}
	word := greeter.Greet()
	fmt.Println(word)
	wordDefault := greeter.GreetDefaultMsg()
	fmt.Println(wordDefault)

}
