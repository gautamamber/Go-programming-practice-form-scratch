// Go’s structs are typed collections of fields. 
// They’re useful for grouping data together to form records.

// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// declare struct

type Employee struct{
	name string
	age int
}

// main function must, run automatically
func main() {
	fmt.Println(Employee{"amber", 20})
	fmt.Println(Employee{name: "Gautam", age:20})
	fmt.Println(Employee{name: "John"})
	fmt.Println(Employee{age:23})
	fmt.Println(Employee{name: "Ann", age: 40})
	s := Employee{name: "Mark", age: 50}
    fmt.Println(s.name)

}
