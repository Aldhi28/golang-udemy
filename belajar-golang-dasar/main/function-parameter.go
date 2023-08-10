package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("hello", firstName,lastName)
}

func main() {
	firstName := "joko"
	sayHello(firstName, "budiman")
	sayHello("Aldhi", "Rizkiyansah")
}