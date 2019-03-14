
// main package
package main 

import "fmt"

// main function m-ust, run automatically
// This is globally defined so it will produce an output
var value1 = "Amber"
// There is a function we called that shorthand but this  we can prevent globally define value to call.
// syntax
// This will raise an error
// name := "Gautam"
func main() {

	// There is a function we called that shorthand but this  we can prevent globally define value to call.
	// Always remember if value is declare but not used it will raise an error in Golang
	// This will print the output
	name := "Gautam"
	age := 34
	price := 11.1
	contact, email := 98765432, "ambergautam1@gmail.com"
	fmt.Println(value1)
	fmt.Println(age)
	fmt.Println(price)
	fmt.Println(name)
	fmt.Println(contact)
	fmt.Println(email)


}

