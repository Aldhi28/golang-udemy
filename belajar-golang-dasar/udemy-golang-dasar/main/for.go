package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// For dengan Statement
	

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
		
	}

	// mengambil data slice dengan mengguanakan for atau perulangan 

	slice := []string {"Aldhi", "Rizki", "Yansah","Budi"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// kede Program for Range 
	for i /*map:key*/, value /*map:value*/ := range slice {
		fmt.Println("index,", i, "value =", value) // map key dan value di pasing ke sini
	}

	person := make(map[string]string)
		person["name"] = "Aldhi"
		person["titile"] = "Programer"

		for key, value := range person {
			fmt.Println(key,"=", value)
		}
	
}