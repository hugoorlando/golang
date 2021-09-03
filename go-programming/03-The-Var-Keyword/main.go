package main

import "fmt"

// puedo acceder a la variable z desde cualquier parte de nuestro codigo
var z = 21

func main() {
	num()
	fmt.Println(z)

	var w int
	w = 23
	fmt.Println(w)
}

func num(){
	fmt.Println(z)
}