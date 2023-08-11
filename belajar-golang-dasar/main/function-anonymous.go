package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		// return false
		fmt.Println("You are bloked", name)
	} else {
		// return true
		fmt.Println("Welcome", name)
	}
}
	// cara yang lama untuk membuat validasi
// func blackListAdmin(name string) bool {
// 	return name == "admin"

// }
// func blackListRoot(name string) bool {
// 	return name == "root"
// }

func main() { 
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("Admin", blackList)
	registerUser("Aldhi", blackList)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Aldhi", func(name string) bool {
		return name == "root"
	})
}