package main

import "fmt"

func main() {
	fmt.Println("Hello World!");

	fmt.Println(HelloWorld("Hugo"));
}

func HelloWorld(user_name string) string {
	return fmt.Sprintf("Hi %s", user_name);
}