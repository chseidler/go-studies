package main

import "fmt"

func main() {
	var i int8 = 10
	var j int32
	j = int32(i)
	fmt.Println(j, "\n")

	var f float64 = 200.56
	var d int = int(f)
	fmt.Printf("f is: %.2f\n", f)
	fmt.Printf("d is: %d", d)
}
