package main

import "fmt"

func sayHelloTo(firstName string, middleName string, lastName string) {
	fmt.Println("Hello,", firstName, middleName, lastName)
}

func main() {
	sayHelloTo("Dyah", "Lutfia", "Aziza")
	sayHelloTo("Ibnu", "Ahmad", "Fauzi")
}
