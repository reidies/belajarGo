package main

import "fmt"

func random () interface{} {
	return "something"
}

func switchingexampe(){

	var result = random()
	switch value := result.(type){
	case string:
		fmt.Println("this", value, "is string")
	case int:
		fmt.Println("this", value, "is int")
	case bool:
		fmt.Println("this", value, "is bool")
	default:
		fmt.Println("somethign else")
	}


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