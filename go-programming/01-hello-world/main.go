package main

import "fmt"

func main() {
	word := "World"
	fmt.Println(greeting(word))
}

func greeting(name string) string {
	return fmt.Sprint("Hello ", name)
}