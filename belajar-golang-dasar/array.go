package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "eko"
	names[1] = "kurniawan"
	names[2] = "khannedy"

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values)

	var exluded = [...]string{
		"here",
		"now",
	}

	fmt.Println(exluded)

	fmt.Println(len(exluded))
}
