package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	aldhi := Man{"Aldhi"}
	aldhi.Married()

	fmt.Println(aldhi.Name)
}