package main

import "fmt"

var a int
var b string = "Programa"
var c bool
var d bool = true

func main() {
	e := 42
	f := "dice hola mundo"
	g := `El doctor dice que comer vegetales es saludable`

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Mi", f, g)
	fmt.Println(e)

	// Printf
	// Verbo %v para imprimir el valor de la variable
	fmt.Printf("El valor de la variable a es: %v\n", a)
	fmt.Printf("El valor de la variable b es: %v\n", b)
	fmt.Printf("El valor de la variable c es: %v\n", c)
	fmt.Printf("El valor de la variable d es: %v\n", d)

	// Verbo %T para imprimir el tipo de la variable
	fmt.Printf("El valor de la variable a es: %T\n", a)
	fmt.Printf("El valor de la variable b es: %T\n", b)
	fmt.Printf("El valor de la variable c es: %T\n", c)
	fmt.Printf("El valor de la variable d es: %T\n", d)

	// Verbo %s para imprimir el valor de los strings
	fmt.Printf("El valor de la variable a es: %s\n", b)
	
	
}