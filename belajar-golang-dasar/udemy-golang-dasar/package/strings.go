package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aldhi Rizkiyansah", "Aldhi")) // mengecek kalimat hasil boolean hasil true
	fmt.Println(strings.Contains("Aldhi Rizkiyansah", "Budi")) // mengecek kalimat hasil boolean hasil false

	fmt.Println(strings.Split("Aldhi Rizkiyansah", " "))

	fmt.Println(strings.ToLower("Aldhi Rizkiyansah"))
	fmt.Println(strings.ToUpper("Aldhi Rizkiyansah"))
	fmt.Println(strings.Title("Aldhi Rizkiyansah"))

	fmt.Println(strings.Trim("   Aldhi Rizkiyansah    "," "))

	fmt.Println(strings.ReplaceAll("eko", "eko","Aldhi"))
}