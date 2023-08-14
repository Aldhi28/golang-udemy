package main

import "fmt"

func main() {
	var name = "Aldhi"

	if name == "Aldhi" {
		fmt.Println("Hello Aldhi")
	}else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Aldhi Rizkiyansah" {
		fmt.Println("Hello joko")
	}else {
		fmt.Println("Hello Aldhi Rizkiyansah")
	}

	// kode Program if Short Statement

	// cara lama untuk menuntukan panjang string
	var length = len(name)
	if length > 5 {
		fmt.Println("Name terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// cara baru untuk menuntukan panjang string
	
	if length := len(name); length > 5 {
		fmt.Println("Name terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}


























}