package main

import "fmt"

// Type Declarations
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ibnu", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
