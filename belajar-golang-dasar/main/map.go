package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Aldhi",
		"address": "Jakarta",
	}

	person["title"] = "Programer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["Title"] = "Belajar Golang"
	book["Author"] = "Aldhi"
	book["Ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Ups")

	fmt.Println(book)
	fmt.Println(len(book))







}