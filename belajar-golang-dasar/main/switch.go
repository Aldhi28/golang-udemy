package main

import (
	"fmt"
)

func main() {
	name := "Aldhi"

	switch name {
	case "Aldhi":
		fmt.Println("Hello Aldhi")
		fmt.Println("Hello Aldhi")
	
	case "Budi":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("kenalan dong")
		fmt.Println("kenalan dong")
	}

	// Switch dengan Short Statement
	// switch length := len(name); length > 5 {
	// 	case true:
	// 		fmt.Println("Name terlalu panjang")
	// 	case false:
	// 		fmt.Println("Nama sudah benar")
	// }

	// Kode Program Switch Tanda Kondisi
	length := len(name)
	switch {
	case length > 15:
		fmt.Println("Name terlalu panjang")
	case length > 5:
		fmt.Println("Name terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	
	}

}