package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye" + name
}

func main() {
	sayGoodBte := getGoodBye

	result := sayGoodBte("Aldhi")

	fmt.Println(result)
}