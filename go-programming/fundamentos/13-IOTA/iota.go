package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	// IOTA nos permite hacer una declaracion de constantes sucesivas y que lo va incrementando en una unidad
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
