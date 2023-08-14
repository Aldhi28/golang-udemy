package main

import "fmt"

type Customer struct {
	Name, Addres string
	Age          int
}

// struct method & struct function
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello",name, "My Name is", customer.Name)
}

func (a Customer) sayHuuuu() {
	fmt.Println("huuuuu from", a.Name)
}


func main() {
	var aldhi Customer
	aldhi.Name = "Aldhi"
	aldhi.Addres = "Lebajak Bulus"
	aldhi.Age = 26

	aldhi.sayHi("Budi")
	aldhi.sayHuuuu()

	// fmt.Println(aldhi)
	// fmt.Println(aldhi.Name)
	// fmt.Println(aldhi.Addres)
	// fmt.Println(aldhi.Age)

	// joko := Customer {
	// 	Name: "joko",
	// 	Addres: "Cirebon",
	// 	Age: 32, 
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Jakarta", 32}
	// fmt.Println(budi)




}