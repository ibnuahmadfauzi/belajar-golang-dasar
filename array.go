package main

import "fmt"

func main() {
	// Deklarasi array
	var names [4]string

	// Inisialisasi array
	names[0] = "Dyah Lutfia Aziza"
	names[1] = "Ibnu Ahmad Fauzi"
	names[2] = "Megumin"
	names[3] = "Kazuma"

	// Modifikasi nilai array
	names[3] = "Aqua"
	names[3] = ""
	names[2] = ""

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	// Deklarasi + inisialisasi
	var values = [...]int{
		90,
		100,
		75,
		300,
		123,
	}

	values[2] = 0
	fmt.Println(values, len(values), "data")
}
