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
	
}