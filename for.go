package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Dyah", "Lutfia", "Aziza"}
	for _, name := range names {
		fmt.Print(name + " ")
	}
}
