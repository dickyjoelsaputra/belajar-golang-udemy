package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpDicky NoKTP = "12321312321"
	var Nikah Married = false
	fmt.Println(ktpDicky)
	fmt.Println(Nikah)
	fmt.Println(NoKTP("0000000000000"))
}
