package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Brooo"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Aldhi")
	fmt.Println(result)

	fmt.Println(getHello(""))
}