package main

import (
	"fmt"
	"strconv"
)

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
