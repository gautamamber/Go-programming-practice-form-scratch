
// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

func addition(num1 int, num2 int) int {
	return num1 + num2
}

func subtraction(num1 int, num2 int) int {
	return num1 - num2
}
func multiplication(num1 int, num2 int) int {
	return num1 * num2
}
func division(num1 int, num2 int) int {
	return num1 % num2
}

// main function must, run automatically
func main() {
	fmt.Println("Hello world")
	fmt.Println(addition(32,20))
	fmt.Println(subtraction(32,20))
	fmt.Println(multiplication(32,20))
	fmt.Println(division(32,20))

}

// save and run command   "go run main.go"