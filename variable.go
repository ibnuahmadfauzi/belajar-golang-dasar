package main

import "fmt"

func main() {
	// Deklarasi + inisialisasi variabel
	name := "Dyah Lutfia Aziza"

	// Output
	fmt.Println("Nilai name :", name)

	// Mengubah nilai variabel
	name = "Ibnu Ahmad Fauzi"
	fmt.Println("Nilai name :", name)

	// Multiple deklarasi variabel
	var (
		gf = "Fia"
		bf = "Ibnu"
	)
	fmt.Println("Couple :", gf, "<3", bf)
}
