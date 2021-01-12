package main

import "fmt"

func main() {

	fmt.Print("hello")

	variables()
	conditional()
	arrays()
	mapping()
}


func variables(){
	var (
		firstName = "heman"
		lastName = "setia"
	)

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
	books ["harga"] = string(123)
	fmt.Println(books)
	fmt.Println(books["title"], books["title"] ,books["harga"])
}
