package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAppLIcation(value int) {
	defer logging() // fungsi defer running compil walaupun salah
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runAppLIcation(10)
}