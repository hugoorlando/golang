// package main declaracion del nombre del paquete
package main

// en import se coloca todos los paquetes que se importan
// cuando se importan multiples paquetes se utilizan parentesis ()
import "fmt"

// func main es el punto de inicio del programa
// en main se llama al paquete importado
func main() {

	word := "World"
	fmt.Println(greeting(username))

	// paquete fmt - funcion Println
	fmt.Println(greeting(word))
}


func greeting(name string) string {
	return fmt.Sprint("Hello ", name)
}