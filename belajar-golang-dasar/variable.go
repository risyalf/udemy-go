package main

import "fmt"

func main() {
	// var name = "Eko Kurniawan"
	// fmt.Println(name)

	// name = "Eko Khannedy"
	// fmt.Println(name)

	name := "Eko Kurniawan" // cuma deklarasi pertama
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	var (
		firstName  = "Eko"
		middleName = "Kurniawan"
		lastName   = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
