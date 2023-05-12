package main

import "fmt"

func main() {
	var name string

	name = "dicky"
	fmt.Println(name)
	name = "joel"
	fmt.Println(name)

	var name2 = "dicky"
	fmt.Println(name2)
	name2 = "joel"
	fmt.Println(name2)

	var age int8 = 30
	fmt.Println(age)

	name3 := "joel"
	fmt.Println(name3)
	name3 = "miami"
	fmt.Println(name3)

	var (
		firstname = "dicky"
		lastname  = "joel"
	)
	fmt.Println(firstname, lastname)
}
