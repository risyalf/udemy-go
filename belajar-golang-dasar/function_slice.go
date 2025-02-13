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

	slice := names[2:]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	append := append(slice, "risyal")

	fmt.Println(append)
}
