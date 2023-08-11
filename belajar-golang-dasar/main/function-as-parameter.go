package main

import "fmt"

// Function Type Declarations
type Filter func(string) string
func sayHelloWithFilter(name string, filter Filter) {
	nameFilterd := filter(name)
	fmt.Println("Hello", nameFilterd)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return "Hello " + name
	}
}





func main() {
	sayHelloWithFilter("Aldhi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
