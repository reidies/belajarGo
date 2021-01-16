package main

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

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
