package main

import "fmt"

type Gf struct {
	Name string
}

func (gf *Gf) People() {
	gf.Name = "Miss. " + gf.Name
}

func main() {
	fia := Gf{"Dyah Lutfia Aziza"}
	fia.People()

	fmt.Println(fia.Name)
}
