package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 81

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi > 80

	lulus := lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}
