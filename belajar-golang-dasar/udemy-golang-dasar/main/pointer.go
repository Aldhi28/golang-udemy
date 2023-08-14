package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// ini untuk cara complex 
	// var address1 Address = Address{"subang", "jawa Barat", "Indonesia"}
	// var address2 *Address = &address1
	// var address3 *Address = &address1
	 /** tanda *Address dan &address adalah untuk menujukan bahwa dia adalah pointer atau reference ke struck sebelumnya dan untuk tanda & dari &address adalah pointer dari variable sebelumnya yang di gunakan untuk merubah data di valuenya*/

	// cara yang sederhana
	// address1 := Address{"subang", "jawa Barat", "Indonesia"}
	// address2 := &address1

	var address1 Address = Address{"subang", "jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "jawa Timur", "Indonesia"}

	/**
	pointer ini *address2 memaksa untuk menggati value dari Address{"subang", "jawa Barat", "Indonesia"}
	*/

	fmt.Println(address1) // {Malang jawa Timur Indonesia}
	fmt.Println(address2) // &{Malang jawa Timur Indonesia}
	fmt.Println(address3) // &{Malang jawa Timur Indonesia}

	// kode program Function new
	var address4 *Address = new(Address)
	address4.City = "jakarta"
	fmt.Println(address4) // &{jakarta  }

	var alamat = Address {
		City: "subang",
		Province: "Jawa Barat",
		Country: "",
	} 
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)


}