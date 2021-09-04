package main

import "fmt"

type dinero int 
var b dinero

func main() {

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

