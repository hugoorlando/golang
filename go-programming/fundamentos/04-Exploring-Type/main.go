package main

import "fmt"

var a string
var b int

func main() {
	// cuando hacemos declaraciones de variables debemos decirle al compilador cual es el tipo de dato que esa variable va a utilizar
	// declarar una variable de cierto tipo solo puede contener valores de ese tipo
	a = "John"
	fmt.Println(a)

	b = 27
	fmt.Println(b)

	var c string
	c = "Brasil"

	// Esto no se puede hacer ya que el tipo string a no puede recibir un int y el tipo int b no puede recibir un tipo string
	// a = 27
	// b = "John"
	// fmt.Println(a,b)

}