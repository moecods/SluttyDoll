package main

import "fmt"

type Person struct {
	Name string
}

type Employee struct {
	Name string
}

type Speaker interface {
	Speak() string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func (e Employee) Speak() string {
	return "Hello, my name is " + e.Name
}

func main() {
	p := Person{Name: "John"}

	var speakerA Speaker = p
	fmt.Println(speakerA.Speak()) // خروجی: Hello, my name is John

	e := Employee{Name: "Tania"}

	var speakerB Speaker = e
	fmt.Println(speakerB.Speak()) // خروجی: Hello, my name is Tania
}
