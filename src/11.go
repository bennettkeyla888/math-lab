  // Package main

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a random number between 1 and 6
	randomNumber := rand.Intn(6) + 1

	// Print the result
	fmt.Println("Random number:", randomNumber)
}