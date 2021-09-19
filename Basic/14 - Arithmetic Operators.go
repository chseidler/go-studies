package main

import "fmt"

func main() {
	fmt.Println(1 + 6)
	fmt.Println(1 - 6)
	fmt.Println(1 * 6)
	fmt.Println(1 / 6)
	fmt.Println(1 % 6)

	p := 20
	d := 15

	c := p + d
	fmt.Printf("The result of addition is: %d \n", c)

	c1 := p - d
	fmt.Printf("The result of subtraction is: %d \n", c1)

	c2 := p * d
	fmt.Printf("The result of multiplication is: %d \n", c2)

	c3 := p / d
	fmt.Printf("The result of division is: %d \n", c3)

	c4 := p % d
	fmt.Printf("The result of modulus is: %d \n", c4)
}
