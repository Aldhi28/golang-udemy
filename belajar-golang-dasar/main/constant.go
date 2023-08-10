package main

import "fmt"

func main() {

	// deklarasi constant biasa
	// const firstName string = "Aldhi"
	// const lastName string = "Rizkiyansah"
	// const value = 1000

	//deklarasi multipel constant 
	const (
		firstName string = "Aldhi"
		lastName string = "Rizkiyansah"
		value = 1000

	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}