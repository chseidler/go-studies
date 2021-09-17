package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "this is a string package"
	b := "THIS IS A STRING PACKAGE UPPER CASE"

	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToLower(b))
}
