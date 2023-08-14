package main

import (
	"fmt"
	"udemy-golang-dasar/database"
)

// _ fungsi untuk ignore

func main() { 
	result := database.GetDatabase()
	fmt.Println(result)
}