package main

import "fmt"

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

	newCustomer := Customer{
		Name:    firstName,
		Address: "tangerang",
		age:     200,
	}

	newCustomer.greetings(firstName)

	fmt.Println(newCustomer.Name, newCustomer.Address, newCustomer.age)

	fmt.Println(firstName, lastName, age, job, marriedStatus)
	fmt.Println(nilai64, nilai8, e, eString)

}