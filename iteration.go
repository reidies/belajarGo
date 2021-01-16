package main

import "fmt"

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
