package main

import "fmt"

type Address struct{
	city, province, country string
}

func (address *Address) changeCity(){
	address.city = "NEW" + address.city
}

func pointer() {
	var address1 = Address{"bandung", "jawabarat", "indonesia"}
	var address2 *Address = &address1

	address2.city = "cirebon"

	fmt.Println(address1, address2)
	address2 = &Address{"Jakarta", "DKI", "Indonesia"}
	fmt.Println(address1,address2)

	*address2 = Address{"Sydney", "NSW", "Australia"}
	fmt.Println(address1,address2)

	alamat := new(Address)
	alamat.city = "sf"
	fmt.Println(alamat)
	changeAddress(alamat)
	fmt.Println(alamat)

	alamat.changeCity()
	fmt.Println(alamat)

}

func understandingPointer(){
	a:=42
	fmt.Println(a)
	fmt.Println(&a)

	b := &a
	fmt.Println(b)
	fmt.Println(&b)
	*b = 43
	fmt.Println(b)
	fmt.Println(&b)

	c := a
	fmt.Println(c)
	fmt.Println(&c)
	c = 45
	fmt.Println(c)
	fmt.Println(&c)
}

func changeAddress (address *Address){
	address.country ="USA"
}
