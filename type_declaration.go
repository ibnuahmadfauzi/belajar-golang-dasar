package main

import "fmt"

func main() {
	// Membuat tipe data baru
	type NIK string

	// Menggunakan tipe data baru
	var ktpIbnu NIK = "3505171600002"

	// Conversion ke tipe data NIK
	var ktpRandom string = "3505171601113"
	var ktpRandomFiks NIK = NIK(ktpRandom)

	// Output
	fmt.Println("NIK Ibnu     :", ktpIbnu)
	fmt.Println("NIK Random   :", ktpRandomFiks)
}
