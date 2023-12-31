package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Nama anda di blacklist")
	} else {
		fmt.Println("Akun berhasil dibuat")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Ibnu", blacklist)
	registerUser("Anjing", func(name string) bool { return name == "Anjing" })
}
