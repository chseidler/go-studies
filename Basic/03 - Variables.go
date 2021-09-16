package main

import "fmt"

func main() {
	var i int
	fmt.Println("Variable i is:", i)
	i = 25
	fmt.Println("Variable i is:", i)

	var i1 int = 25
	fmt.Println("\nVariable i1 is:", i1, "\n")

	i2 := 569875632
	fmt.Println(i2)

	var i3, j int = 25, 50
	fmt.Println(i3, j)

	var a float64 = 3.14
	fmt.Println(a)

	var b = true
	fmt.Println(b)
}
