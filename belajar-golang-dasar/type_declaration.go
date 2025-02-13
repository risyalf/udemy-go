package main

import "fmt"

func main() {
	type NoKTP string

	var noKtpEko NoKTP = "12345"

	var contoh = "112233"

	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(noKtpEko)
	fmt.Println(contohKtp)
}
