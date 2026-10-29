package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Bow Bow")
}

func main() {
	var a Animal

	a = Dog{}

	a.Speak()
}