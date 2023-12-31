package main

import "fmt"

type Mhs struct {
	Name, Address string
	Age           int
}

func (mhs Mhs) dataOutput() {
	fmt.Println("=============================")
	fmt.Println("Name    :", mhs.Name)
	fmt.Println("Address :", mhs.Address)
	fmt.Println("Age     :", mhs.Age)
	fmt.Println("=============================")
}

func main() {
	var mhs1 Mhs
	mhs1.Name = "Dyah Lutfia Aziza"
	mhs1.Address = "Kota Blitar"
	mhs1.Age = 19
	mhs1.dataOutput()

	// Struct Literals
	mhs2 := Mhs{
		Name:    "Ibnu Ahmad Fauzi",
		Address: "Kabupaten Blitar",
		Age:     23,
	}
	mhs2.dataOutput()
}
