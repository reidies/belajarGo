package main

import "fmt"

type abstracts struct {
	T1 string
	T2 int
	T3 bool
	T4 float64
}

type Customer struct{
	Name, Address string
	age int
}

type istCorrect interface {
	validate(string) bool
}

func (customer Customer)validate(input string) bool{
	return input == "herman"
}


func (customer Customer) greetings(name string){
	fmt.Println("hello", name)
}

//interface
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

//implementation
func (person Person) GetName() string {
	return person.Name
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

