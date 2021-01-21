package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)


func main() {

	//fmt.Println("1+2 :",mySum(1,2))
	//values :=[] int{5,1,2,5,1}
	//maximumValue(values...)

	concurrency()
	//understandingPointer()

}

func jsonMarshal(){
	c1 := Customer{
		Name: "rudi",
		Address: "sampit",
		age: 32,
	}
	c2 := Customer{
		Name: "alex",
		Address: "sumatera",
		age: 42,
	}

	customers := []Customer{c1,c2}

	bs, err := json.Marshal(customers)
	if err != nil{
		log.Println("something wrong", err)
	}else {
		fmt.Println(string(bs))
	}

}

func mySum(xi ...int) int{
	sum:=0
	for _, v:=range xi{
		sum +=v
	}
	return sum
}


func openFile(){
	//_, err := os.OpenFile("filename.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	_,err := os.Open("filename.txt")
	if err!=nil{
		//fmt.Println(err)
		log.Fatal(err)
		//log.Panic(err)
	}
}

func simpleGolang(){
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








