package main

import "fmt"

// Declarando la variable con el identificador z de tipo int
var z int
var a bool

func main() {
	z = 21
	fmt.Println(z)

	// Valor cero para: 
	// booleanos, integers, floats, "" strings vacios

	// nil para:
	// pointers, funciones, interfaces, slices, channels, maps

	// var se usa cuando inicializamos variables con valores 0
	fmt.Println(a)
}