package main
import "fmt"
func main() {
	a := 5
	b := 10
	c := add(a, b)
	fmt.Println("The sum is", c)
}
func add(a int, b int) int {
	return a + b
}