package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 0 {
			break
		}
		fmt.Println("Perulanggan ke",i)
	}
}