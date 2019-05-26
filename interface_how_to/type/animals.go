package _type

//This is a core concept in Go’s
//type system; instead of designing our abstractions in terms of what kind of data our types can hold,
// we design our abstractions in terms of what actions our types can execute.
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Wof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

type Llama struct{}

func (l Llama) Speak() string {
	return "????"
}

type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string {
	return "Love Java"
}

type Anything interface {
}

type Apple struct{}

func Eat(anything Anything) string {
	return "Eat le"
}
