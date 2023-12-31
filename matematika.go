package main

import "fmt"

func main() {
	// Aritmatika
	var a = 10
	var b = 20
	var d = 5
	var c = a + b*d // 110

	// Augmented assigements
	c += 40 // 150

	// Unary operator
	c++ // 151
	c-- // 150

	// Output
	fmt.Println(c)
}
