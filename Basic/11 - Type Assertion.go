package main

import "fmt"

type Animal interface {
	call() string
}

type dog struct {
	name string
}

type cow struct {
	name string
}

func (d *dog) call() string {
	return fmt.Sprintf("Au Au Au %s Woof", d.name)
}

func (c *cow) call() string {
	return fmt.Sprintf("Muuuu %s Muuuu", c.name)
}

func main() {
	var animals []Animal = []Animal{
		&dog{"German"},
		&cow{"Australian"},
	}
	for _, animal := range animals {
		fmt.Println(animal.call())
	}
}
