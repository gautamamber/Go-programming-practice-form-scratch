// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// main function must, run automatically
func main() {

	// Arrays are the collection of elements of similar data types
	// No of elements in arrays are fixed

	// Declare
	var totalValue [4]int
	totalValue[0] = 1
	totalValue[1] = 2
	totalValue[2] = 3
	totalValue[3] = 4

	fmt.Println(totalValue)
	fmt.Println(totalValue[0])
	fmt.Println(totalValue[2])
	fmt.Println(totalValue[1])
	fmt.Println(totalValue[3])

	// we can also declare array using shorthand functions

	newValues := [3]int{1,2,3}
	fmt.Println(newValues)
}

// save and run command   "go run main.go"