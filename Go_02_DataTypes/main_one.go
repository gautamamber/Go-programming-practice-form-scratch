
// main package
package main 

import "fmt"

// main function must, run automatically
func main() {

	// If we use const then it is not possible to modify a value , it will raise an error 

	const value1 = false

	// value1 = true   (It will raise an error)

	// Always remember if value is declare but not used it will raise an error in Golang

	fmt.Println(value1)


}
