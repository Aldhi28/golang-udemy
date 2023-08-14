package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a[a-z]i") 

	fmt.Println(regex.MatchString("aldi"))
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("afi"))
	fmt.Println(regex.MatchString("aGi"))

	search := regex.FindAllString("aldhi adi aki ali ani", -1)
	fmt.Println(search)



}