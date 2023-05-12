package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	fmt.Println(a + b)

	a = a * 10
	fmt.Println(a)

	b += 10
	fmt.Println(b)
}
