package main

import "fmt"

func main() {
	var names = [...]string{
		"Eko",
		"Kurniawan",
		"Khannedy",
		"Budy",
		"Nugraha",
	}

	slice1 := names[4:5]
	fmt.Println(slice1)

	slice1 = names[1:]
	fmt.Println(slice1)

	slice1 = names[:2]
	fmt.Println(slice1)

	slice1 = names[:]
	fmt.Println(slice1)
}
