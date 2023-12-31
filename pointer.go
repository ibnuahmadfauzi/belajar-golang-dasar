package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer

	address2.City = "Blitar"
	address2.Province = "Jawa Timur"

	fmt.Println(address1)
	fmt.Println(*address2)
}
