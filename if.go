package main

import "fmt"

func main() {
	name := "Fia"

	if name == "Fia" {
		fmt.Println("Nama benar adalah Fia")
	} else if name == "Ibnu" {
		fmt.Println("Nama benar adalah Ibnu")
	} else {
		fmt.Println("Nama salah bukan Fia ataupun Ibnu")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah sesuai")
	}
}
