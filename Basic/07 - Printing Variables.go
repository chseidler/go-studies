package main

import "fmt"

var (
	a = 600
	b = false
	c = 3.14
	d = "Go"
)

func main() {
	fmt.Printf("The value of a is: %d \n", a)
	fmt.Printf("The value of b is: %t \n", b)
	fmt.Printf("The value of c is: %g \n", c)
	fmt.Printf("The value of d is: %s \n", d)
}
