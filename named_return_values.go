package main

import "fmt"

func getCoupleName() (gf, bf string) {
	gf = "Dyah Lutfia Aziza"
	bf = "Ibnu Ahmad Fauzi"
	return gf, bf
}

func main() {
	gf, bf := getCoupleName()
	fmt.Println("GF   :", gf)
	fmt.Println("BF   :", bf)
}
