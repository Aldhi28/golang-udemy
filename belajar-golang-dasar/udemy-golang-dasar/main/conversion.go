package main

import "fmt"

func main() {

	// konversi tipe data int 
	var nilai32 int32 = 100000 // 10000
	var nilai64 int64 = int64(nilai32) //10000
	var nilai8 int8 = int8(nilai32) //-96

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Aldhi"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}