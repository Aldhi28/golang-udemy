package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Aldhi")
	data.PushBack("Rizki")
	data.PushBack("Yansah")
	data.PushFront("Budi")


	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)


    // dari depan ke belakang
	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
	