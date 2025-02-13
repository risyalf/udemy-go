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

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]

	fmt.Println(daySlice1)
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"

	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println(days)
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "eko"
	newSlice[1] = "eko"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "budi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "risyal"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	newSlice3 := append(newSlice2, "budi2")
	newSlice4 := append(newSlice3, "budi3")
	newSlice5 := append(newSlice4, "budi4")

	newSlice5[0] = "hehe"
	newSlice4[0] = "hehehe"
	fmt.Println(newSlice5)
	fmt.Println(newSlice)

	fromSlize := days[:]
	toSlice := make([]string, len(fromSlize), cap(fromSlize))

	copy(toSlice, fromSlize)
	fmt.Println(fromSlize)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
