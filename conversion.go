package main

import "fmt"

func main() {
	// Conversion antar int
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)
	var nilai8 int8 = int8(nilai16)

	// Conversi number string
	var name string = "Dyah Lutfia Aziza"
	var e uint8 = name[3]
	var eString = string(e)

	fmt.Println("Nilai32 :", nilai32)
	fmt.Println("Nilai64 :", nilai64)
	fmt.Println("Nilai16 :", nilai16)
	fmt.Println("Nilai8  :", nilai8)
	fmt.Println("String  :", eString)
}
