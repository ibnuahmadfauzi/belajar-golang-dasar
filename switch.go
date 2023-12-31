package main

import "fmt"

func main() {
	name := "Fia"

	switch name {
	case "Fia":
		fmt.Println("Pacar saya")
	case "Dyah":
		fmt.Println("Pacar saya")
	case "Ibnu":
		fmt.Println("Nama saya")
	default:
		fmt.Println("Kamu siapa ?")
	}

	// Shorthand
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah valid")
	}

	// Switch tanpa kodisi
	length := len(name)
	switch {
	case length > 2:
		fmt.Println("Nama terlalu panjang")
	case length <= 2:
		fmt.Println("Nama sudah valid")
	}
}
