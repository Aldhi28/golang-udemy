package main

import (
	"fmt"
)

func endApp() {
	message := recover() // fungsi untuk mejalankan setelah fungsi panic di jalankan dan untuk menangkap message dari value panic
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp () // untuk menjalankan aplikasi walaupun error
		if error {
			panic("APLIKASI ERROR") // untuk menghentikan aplikasi
		}
		fmt.Println("Aplikasi berjalan")
	
}

func main() {
	runApp(true)// untuk menentukan running aplikasi true or false 
	fmt.Println("Aldhi")
}