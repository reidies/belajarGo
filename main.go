package main

import (
	"fmt"
	"strconv"
)

type abstracts struct {
	T1 string
	T2 int
	T3 bool
	T4 float64
}

func main() {

	fmt.Print("hello")

	variables()
	conditional()
	arrays()
	mapping()
	switchingexampe()
	iteration()

	value1:= 20
	value2:=21

	fmt.Println("calculating", calculation(value1,value2))
	bigger,smaler := isItbigger(value1,value2)
	fmt.Println("is it bigger", bigger, "istItsmaller", smaler)
	multiple, division := multipleDivision(float64(value1), float64(value2))
	fmt.Println("multiple", multiple, "division", division)

	numbers := []int {10,20,30,40,50}
	fmt.Println("the total of", sumAll(numbers...))
}

func calculation(value1 int, value2 int)int{
	//with specific return value
	return value1+value2
}

func isItbigger(value1 int, value2 int) (bool, bool){
	bigger := value1>value2
	smaller := value1<value2
	return bigger, smaller
}

func multipleDivision (value1 float64, value2 float64)(multiplication float64, division float64){
	multiplication = value1 * value2
	division = value2/value1

	return
}

func sumAll(numbers ...int) int{
	total :=0
	for _, number := range numbers {
		total +=number
	}

	return total
}


func variables(){
	var (
		firstName = "heman"
		lastName = "setia"
	)
	sayHelloTo(firstName, lastName)
	age := 32

	const job = "programmer"
	type married bool
	var nilai32 int32 = 1000000
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai32)
	var marriedStatus married = false

	var e = firstName[2] //byte
	var eString = string(e)


	fmt.Println(firstName, lastName, age, job, marriedStatus)
	fmt.Println(nilai64, nilai8, e, eString)

}

func conditional(){
	var name1 = "EKI"
	var name2 = "EKO"

	fmt.Println(name1 == name2, name1<name2)
}

func sayHelloTo(firstName string, lastName string){
	fmt.Println("hello", firstName, lastName)
}

func arrays(){

	var names = [3] string {"adi", "rudi"}
	fmt.Println(names, len(names))
	fmt.Println(names[0], names[1], names[2])
	var months = [...]string{"januari", "februari", "maret", "april", "may", "june", "july", "august", "sepctember",
		"october", "november", "desember"}
	fmt.Println(months[4:7])//slice

	newArray := make([]string, 2, 5)
	newArray[0] = "eko"
	newArray[1] ="budi"

	pizza := []int{1,2,3,4,5} //slice
	fmt.Println(pizza)
	var newPizza = []int{12,3}
	newPizza = append(pizza, newPizza...)
	fmt.Println(newPizza)
}

func mapping(){
	person := map[string]string{
		"names":"asi", "address":"lokasi",
	}
	person["job"] = "programmer"
	fmt.Println(person["names"], person["address"], person["job"])
	var books = make(map[string] string)
	books ["title"] = "golang"
	books ["author"] = "rudi"
	books ["harga"] = strconv.Itoa(123)
	fmt.Println(books)
	fmt.Println(books["title"], books["author"] ,books["harga"])
}

func switchingexampe(){
	name := "budis"

	switch  lenght := len(name); lenght<5 {
	case true :
		fmt.Println("name pendek")
	case false :
		fmt.Println("name panjang")
	}

	length := len(name)
	switch  {
	case length<5:
		fmt.Println("nama pendek")
	case length >5:
		fmt.Println("name panjang")
	default:
		fmt.Println("hello", name)

	}
}

func iteration(){
	for counter:=1; counter<=100;counter++{

		if counter%2 ==0{
			continue
		}
		fmt.Print( counter)

		if counter >10 {break}
	}

	names := []string{"raden", "sakit", "jiwa"}

	for index, name:= range names{
		fmt.Print("index", index, "=", name)
	}

}

