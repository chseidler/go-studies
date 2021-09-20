package main

import "fmt"

func main() {
	a := 20
	a += 10
	fmt.Println(a)

	var b = 50
	var c = 20

	b = c
	fmt.Println(b)

	b += c
	fmt.Println(b)

	b -= c
	fmt.Println(b)

	b /= c
	fmt.Println(b)

	b %= c
	fmt.Println(b)

	b *= c
	fmt.Println(b)
}
