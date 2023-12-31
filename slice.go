package main

import "fmt"

func main() {
	months := [...]string{
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

	choiceMonths := months[5:8]

	// fmt.Println("array    :", months)
	fmt.Println("data         :", choiceMonths)
	fmt.Println("length       :", len(choiceMonths))
	fmt.Println("capacity     :", cap(choiceMonths))
	choiceMonths = append(choiceMonths, "Data Baru")
	choiceMonths = append(choiceMonths, "Data Baru Lagi")
	fmt.Println("new data     :", choiceMonths)
	fmt.Println("new length   :", len(choiceMonths))
	fmt.Println("new capacity :", cap(choiceMonths))

	// fmt.Println("full data    :", months)

	// membuat slice sendiri
	newSlice := []string{
		"Dyah Lutfia Aziza",
		"Ibnu Ahmad Fauzi",
	}
	fmt.Println("new slice    :", newSlice)
	fmt.Println("length       :", len(newSlice))
	fmt.Println("capacity     :", cap(newSlice))

	newSlice = append(newSlice, "Kazuma")
	newSlice = append(newSlice, "Megumin")
	fmt.Println("new slice    :", newSlice)
	fmt.Println("length       :", len(newSlice))
	fmt.Println("capacity     :", cap(newSlice))

	// copy slice
	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("copy slice   :", toSlice)
}
