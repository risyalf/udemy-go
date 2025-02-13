package main

import "fmt"

func main() {
	name := "ooaoos"

	if name == "eko" {
		fmt.Println("it is eko!")
	} else if name == "budi" {
		fmt.Println("it is budi!")
	} else {
		fmt.Println("it's not eko nor budi!")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("panjang")
	} else {
		fmt.Println("pendek")
	}
}
