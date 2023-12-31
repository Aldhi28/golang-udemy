package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println("Argument :",args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	}else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP-USERNAME")
	password := os.Getenv("APP-PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}