package main

import "fmt"

func main() {
	// create map
	mhs1 := map[string]string{
		"name":     "Dyah Lutfia Aziza",
		"prodi":    "Hukum",
		"fakultas": "Ilmu Hukum",
	}

	fmt.Println("Name        :", mhs1["name"])
	fmt.Println("Prodi       :", mhs1["prodi"])
	fmt.Println("Fakultas    :", mhs1["fakultas"])
	fmt.Println("nim         :", mhs1["nim"])
	fmt.Println("length data :", len(mhs1))

	// delete map
	delete(mhs1, "fakultas")
	fmt.Println("\nUpdate")
	fmt.Println("Name        :", mhs1["name"])
	fmt.Println("Prodi       :", mhs1["prodi"])
	fmt.Println("Fakultas    :", mhs1["fakultas"])
	fmt.Println("nim         :", mhs1["nim"])
	fmt.Println("length data :", len(mhs1))

	// modifikasi map
	mhs1["fakultas"] = "Teknologi Informasi"
	mhs1["name"] = "Ibnu Ahmad Fauzi"
	mhs1["prodi"] = "Teknik Informatika"
	mhs1["nim"] = "19104410028"
	fmt.Println("\nUpdate")
	fmt.Println("Name        :", mhs1["name"])
	fmt.Println("Prodi       :", mhs1["prodi"])
	fmt.Println("Fakultas    :", mhs1["fakultas"])
	fmt.Println("nim         :", mhs1["nim"])
	fmt.Println("length data :", len(mhs1))
}
