package main

import "fmt"

func main() {
	name := "ekos"

	switch name {
	case "eko":
		fmt.Println("it is eko!")
	case "budi":
		fmt.Println("it is budi!")
	default:
		fmt.Println("it is not eko nor budi!")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("panjang")
	case false:
		fmt.Println("pendek")
	}

	lenght := len(name)

	switch {
	case lenght > 5:
		fmt.Println("panjang")
	case lenght < 5:
		fmt.Println("pendek")
	}
}
