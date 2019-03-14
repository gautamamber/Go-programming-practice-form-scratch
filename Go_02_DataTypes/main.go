
// main package
package main 

import "fmt"

// main function must, run automatically
func main() {

	// Types of: Data types
	// string, bool , int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64
	// byte 
	// rune
	// float32, float64
	// complex64, complex128
	// Generally Data types are initialized with var or const

	// Example:

	var firstName string = "Amber"
	//or 
	var lastName = "Gautam" 
	var num1 = 21
	var num2 int = 22
	var num3 float64 = 2.3
	var num4 = 3.2
	var value1 bool = true
	var value2 = false

	// Always remember if value is declare but not used it will raise an error in Golang


	fmt.Println("Hello world")
	fmt.Println("Firstname is " + firstName)
	fmt.Println("Lastname is " + lastName)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(value1)
	fmt.Println(value2)


}

// save and run command   "go run main.go"