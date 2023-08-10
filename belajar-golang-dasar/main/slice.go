package main

import "fmt"

func main() {
	var mounth = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// len (slice) untuk mendapatkan panjang data
	// cap (slice) untuk mendapatkan kapasitas

	var slice1 = mounth[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// mounth[5] = "Diubah" // untuk mengubah data array yang ada di dalam array
	// fmt.Println(slice1) 

	// slice1[0]= "Mei lagi"
	// fmt.Println(mounth) 

	var slice2 = mounth[10:] // [november ,desember]
	fmt.Println(slice2) 

	var slice3 = append(slice2, "Aldhi") // [November Desember Aldhi] 
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember" // [November Bukan Desember Aldhi]
	fmt.Println(slice3)

	fmt.Println(slice2) // [November Desember]
	fmt.Println(mounth) // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember]

	newSLice := make([]string, 2, 5)

	newSLice[0] = "Aldhi"
	newSLice[1] = "Rizkiyansah"

	fmt.Println(newSLice)
	fmt.Println(len(newSLice))
	fmt.Println(cap(newSLice))

	copySlice := make([]string, len(newSLice), cap(newSLice))
	copy(copySlice, newSLice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	


}