package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAdlhi NoKTP = "123456789012"
	var married Married = false
	fmt.Println(noKtpAdlhi)
	fmt.Println(married)
}