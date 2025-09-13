package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (Cat) Speak() string {
	return "Meau!"
}

//

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

// Встраивание (embedding)

type Runner interface {
	Run() error
}

func main() {

	//fmt.Print("Hello world!")

	//animals := []Animal{Dog{}, Cat{}}
	//
	//for index, animal := range animals {
	//	fmt.Println(index, animal.Speak())
	//	//fmt.Println(animal.Speak())
	//}

	var runner Runner

	fmt.Printf("Type: %T\n Value: %v\n", runner, runner)

}
