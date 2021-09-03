package main

import "fmt"

var b bool

func main() {
	// Un booleano es un tipo binario que puede tener dos posibles valores: true o false

	// Operador de comparacion
	// ==, <=, >=, !=, <, >
	
	a := true
	b := false 

	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Printf("El valor b es de tipo: %T\n", b)
	
}