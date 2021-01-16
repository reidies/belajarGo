package main

import (
	"errors"
	"fmt"
	"strings"
)



func main() {

	var eko Person
	eko.Name = "Eko"

	SayHello(eko)

	fmt.Print("hello")

	variables()
	conditional()
	arrays()
	mapping()
	switchingexampe()
	iteration()
	pointer()

	value1:= 20
	value2:=21

	fmt.Println("calculating", calculation(value1,value2))
	bigger,smaler := isItbigger(value1,value2)
	fmt.Println("is it bigger", bigger, "istItsmaller", smaler)
	multiple, division := multipleDivision(float64(value1), float64(value2))
	fmt.Println("multiple", multiple, "division", division)

	numbers := []int {10,20,30,40,50}
	fmt.Println("the total of", sumAll(numbers...))

	fmt.Println(kalimatdifilter("jancok", spamSwear))

	fmt.Println(factorialRecursive(5))

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)

	runApp(true)
}

func endApp(){
	fmt.Println("done")
	message := recover()
	fmt.Println("error message :", message)
}

func runApp(error bool){
	defer endApp()
	if error {
		fmt.Println(errors.New("something error"))
		panic("Error in apps")
	}
}

func kalimatdifilter(kalimat string, filter func(string) string) string {
 	kalimatfilter := filter(kalimat)
	return kalimatfilter
}

func spamSwear(kalimat string) string{
	//swearwords := []string{"jancok", "tai", "anjing"}
	check := strings.Contains(kalimat, "jancok")
	if  check == true{
		return "###"
	}
	return kalimat
}

func sayHelloTo(firstName string, lastName string){
	fmt.Println("hello", firstName, lastName)
}








