package main

import "fmt"

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