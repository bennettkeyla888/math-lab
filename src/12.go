  package main

import "fmt"

func main() {
	const pi = 3.14
	var radius float64 = 5
	area := pi * math.Pow(radius, 2)
	fmt.Println("The area of the circle is: ", area)
}