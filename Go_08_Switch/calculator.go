// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// main function must, run automatically
func main() {
	num1 := 23
	num2 := 33
	case_value := 1
	switch case_value{
	case 1:
		fmt.Println(num1 + num2)
	case 2:
		fmt.Println(num1 - num2)
	case 3:
		fmt.Println(num1 * num2)
	case 4:
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid")
	}
}