package main

import "fmt"

func getFullName() (string, string) {
	return "Aldhi", "Rizkiyansah"
}

func main() {
	firstName, LastName := getFullName()
	fmt.Println(firstName,LastName)
}