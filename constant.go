package main

import "fmt"

func main() {
	// Membuat constant
	const phi = 3.14
	fmt.Println("Nilai phi :", phi)

	// Modifikasi constant
	// phi = 3.15 // ERROR

	// Deklarasi multiple constant
	const (
		gf = "Fia"
		bf = "Ibnu"
	)
	fmt.Println(gf, "&", bf, "together forever")
}
