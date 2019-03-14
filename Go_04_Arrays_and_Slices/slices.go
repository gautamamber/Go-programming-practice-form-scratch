// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// main function must, run automatically
func main() {

	
	// No of Elements in slices are not fixed
	stringValues :=	[]string{"a","b","c"}
	fmt.Println(stringValues)
	fmt.Println(stringValues[1])
	fmt.Println(stringValues[2])
	fmt.Println(stringValues[0])
	
}

// save and run command   "go run main.go"