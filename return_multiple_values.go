package main

import "fmt"

func getFullName() (string, string, string) {
	return "Fia", "Megumin", "Ibnu"
}

func main() {
	gf, _, bf := getFullName()
	fmt.Println(gf, "LOVE", bf)
}
