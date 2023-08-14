package main

import (
	"fmt"
	"udemy-golang-dasar/helper"
)

func main() {
	helper.SayHello("Aldhi")
	// helper.sayGoodbye("Aldhi") //error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //error
}