
// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// In golang main function is run automatically and we can also create our own functions

func fullName(name string) string { //    func keyword <function name with argument> <return type>
	return "Amber" + " " + name
}

// main function must, run automatically
func main() {
	fmt.Println("Hello world")
	fmt.Println(fullName("Gautam"))
}


// save and run command   "go run main.go"